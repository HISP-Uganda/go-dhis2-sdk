package dhis2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HISP-Uganda/go-dhis2-sdk/aggregate"
	"github.com/HISP-Uganda/go-dhis2-sdk/dhis2/schema"
	"github.com/HISP-Uganda/go-dhis2-sdk/tracker"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Client struct {
	BaseURL       string
	Username      string
	Password      string
	Resty         *resty.Client
	MaxRetries    int
	BaseDelay     time.Duration
	RateLimit     time.Duration
	Burst         int
	tokens        chan struct{}
	lastCall      time.Time
	rateInit      sync.Once
	mu            sync.Mutex
	BatchSize     int // Optional: for batching requests
	FailDir       string
	RetentionDays int // Optional: number of days to retain failed batches
}

func NewClient(url, user, pass string) *Client {
	client := resty.New().
		SetBaseURL(url).
		SetBasicAuth(user, pass).
		SetHeader("Content-Type", "application/json")

	return &Client{
		BaseURL:       url,
		Username:      user,
		Password:      pass,
		Resty:         client,
		MaxRetries:    3,
		BaseDelay:     time.Second,
		RateLimit:     0,
		Burst:         1,
		BatchSize:     100,
		FailDir:       "/tmp/dhis2_failures", // Directory to save failed batches
		RetentionDays: 30,                    // Retain failed batches for 30 days
	}
}

// waitForRateSlot enforces rate limiting using a token bucket.
// It blocks if there are no available tokens, and replenishes tokens periodically
// according to the configured RateLimit interval and Burst size.
func (c *Client) waitForRateSlot() {
	if c.RateLimit <= 0 || c.Burst <= 0 {
		return
	}

	c.rateInit.Do(func() {
		c.tokens = make(chan struct{}, c.Burst)
		for i := 0; i < c.Burst; i++ {
			c.tokens <- struct{}{}
		}
		go func() {
			ticker := time.NewTicker(c.RateLimit)
			defer ticker.Stop()
			for range ticker.C {
				select {
				case c.tokens <- struct{}{}:
				default:
				}
			}
		}()
	})
	<-c.tokens
}

func (c *Client) Post(path string, payload any, queryParams ...map[string]string) (*resty.Response, error) {
	r := c.Resty.R().SetBody(payload)
	if len(queryParams) > 0 {
		r.SetQueryParams(queryParams[0])
	}
	return r.Post(path)
}

func (c *Client) Get(path string, result any, queryParams ...map[string]string) (*resty.Response, error) {
	r := c.Resty.R().SetResult(result)
	if len(queryParams) > 0 {
		r.SetQueryParams(queryParams[0])
	}
	return r.Get(path)
}

func (c *Client) Put(path string, payload any, queryParams ...map[string]string) (*resty.Response, error) {
	r := c.Resty.R().SetBody(payload)
	if len(queryParams) > 0 {
		r.SetQueryParams(queryParams[0])
	}
	return r.Put(path)
}

func (c *Client) Delete(path string, queryParams ...map[string]string) (*resty.Response, error) {
	r := c.Resty.R()
	if len(queryParams) > 0 {
		r.SetQueryParams(queryParams[0])
	}
	return r.Delete(path)
}

// SendAggregateDataValues sends a dataValueSets payload to DHIS2 and logs the response.
func (c *Client) SendAggregateDataValues(ctx context.Context, payload *aggregate.DataValueSetPayload) (*aggregate.ImportSummaryResponse, error) {

	if err := payload.Validate(); err != nil {
		log.WithError(err).Error("Validation failed before sending payload")
		return nil, err
	}
	var resp aggregate.ImportSummaryResponse

	res, err := c.Resty.R().
		SetContext(ctx).
		SetBody(payload).
		SetResult(&resp).
		Post("/api/dataValueSets")

	if err != nil {
		log.WithError(err).Error("Failed to send aggregate data to DHIS2")
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	if res.StatusCode() != http.StatusOK {
		log.WithFields(log.Fields{
			"status": res.Status(),
			"body":   res.String(),
		}).Error("DHIS2 returned non-200 status")
		return &resp, fmt.Errorf("dhis2 error response: %s", res.Status())
	}

	resp.LogSummary()
	return &resp, nil
}

// SendTrackerPayload sends a nested tracker payload to DHIS2 with optional query parameters.
func (c *Client) SendTrackerPayload(ctx context.Context, payload *tracker.NestedPayload, queryParams map[string]string) (*tracker.TrackerResponse, *resty.Response, error) {
	var res *resty.Response
	var err error
	var importRes schema.TrackerImportReport
	var asyncRes tracker.AsyncResponse

	isAsync := strings.EqualFold(queryParams["async"], "true")

	for attempt := 1; attempt <= c.MaxRetries; attempt++ {
		c.waitForRateSlot()

		req := c.Resty.R().
			SetContext(ctx).
			SetBody(payload).
			SetHeader("Content-Type", "application/json")

		for key, val := range queryParams {
			req.SetQueryParam(key, val)
		}

		res, err = req.Post("/api/tracker")
		if err != nil {
			log.WithError(err).WithField("attempt", attempt).Error("Failed to send tracker payload")
		} else if res.StatusCode() != http.StatusOK && res.StatusCode() != http.StatusCreated {
			log.WithFields(log.Fields{
				"status":  res.StatusCode(),
				"body":    res.String(),
				"attempt": attempt,
			}).Error("DHIS2 tracker payload rejected")
		} else {
			if isAsync {
				if err := json.Unmarshal(res.Body(), &asyncRes); err != nil {
					return nil, res, fmt.Errorf("failed to parse async response: %w", err)
				}
				jobID := ""
				if asyncRes.Response != nil && asyncRes.Response.ID != "" {
					jobID = asyncRes.Response.ID
				} else if asyncRes.ID != nil {
					jobID = *asyncRes.ID
				}
				log.WithField("jobID", jobID).Info("Async tracker import submitted")
				return &tracker.TrackerResponse{AsyncResp: &asyncRes}, res, nil
			} else {
				if err := json.Unmarshal(res.Body(), &importRes); err != nil {
					log.WithError(err).Error("Failed to parse tracker import response")
					return nil, res, fmt.Errorf("failed to parse response: %w", err)
				}
				if validationReport := importRes.ValidationReport; validationReport != nil {
					if len(validationReport.ErrorReports) > 0 {
						errBytes, _ := json.MarshalIndent(validationReport.ErrorReports, "", "  ")
						log.WithField("errors", string(errBytes)).Warn("Validation errors reported by DHIS2")
						return &tracker.TrackerResponse{SyncResp: &importRes}, res, errors.New("tracker import failed with validation errors")
					}
					if len(validationReport.WarningReports) > 0 {
						warnBytes, _ := json.MarshalIndent(validationReport.WarningReports, "", "  ")
						log.WithField("warnings", string(warnBytes)).Info("Validation warnings reported by DHIS2")
					}
				}
				log.WithField("status", res.Status()).Info("Tracker payload submitted successfully")
				return &tracker.TrackerResponse{SyncResp: &importRes}, res, nil
			}
		}

		if attempt < c.MaxRetries {
			delay := c.BaseDelay * time.Duration(1<<uint(attempt-1))
			log.WithField("delay", delay).Infof("Retrying in %v...", delay)
			time.Sleep(delay)
		}
	}

	return nil, res, fmt.Errorf("failed to submit payload after %d attempts: %w", c.MaxRetries, err)
}

// SendBatchedTrackerPayload sends a large tracker payload in batches using SendTrackerPayload.
// It collects all successful responses and continues through retries and failures.
// If a batch fails, its payload is written to disk for later inspection.
// SendBatchedTrackerPayload splits and sends batches of payloads.
func (c *Client) SendBatchedTrackerPayload(ctx context.Context, payload *tracker.NestedPayload, queryParams map[string]string) ([]*tracker.TrackerResponse, error) {
	var responses []*tracker.TrackerResponse
	var batchErrors []error

	if err := os.MkdirAll(c.FailDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create fail dir: %w", err)
	}

	for start := 0; start < len(payload.TrackedEntities); start += c.BatchSize {
		end := start + c.BatchSize
		if end > len(payload.TrackedEntities) {
			end = len(payload.TrackedEntities)
		}
		subPayload := &tracker.NestedPayload{
			TrackedEntities: payload.TrackedEntities[start:end],
		}
		log.Infof("Sending batch %d to %d of %d tracked entities", start+1, end, len(payload.TrackedEntities))

		resp, _, err := c.SendTrackerPayload(ctx, subPayload, queryParams)
		if err != nil {
			filename := fmt.Sprintf("failed_batch_%d_%d_%s.json", start+1, end, time.Now().Format("20060102_150405"))
			filePath := filepath.Join(c.FailDir, filename)
			data, _ := json.MarshalIndent(subPayload, "", "  ")
			_ = os.WriteFile(filePath, data, 0644)
			log.WithField("file", filePath).Errorf("Batch %d-%d failed, saved to file", start, end)
			batchErrors = append(batchErrors, fmt.Errorf("batch %d–%d failed: %w", start, end, err))
			continue
		}
		responses = append(responses, resp)
	}

	log.WithFields(log.Fields{
		"total":   len(payload.TrackedEntities),
		"batches": (len(payload.TrackedEntities) + c.BatchSize - 1) / c.BatchSize,
		"success": len(responses),
		"failed":  len(batchErrors),
	}).Info("Batch submission summary")

	if len(batchErrors) > 0 {
		return responses, fmt.Errorf("%d batch(es) failed: %v", len(batchErrors), batchErrors)
	}
	return responses, nil
}

// cleanupOldArchivedBatches removes archived batch files older than RetentionDays.
func (c *Client) cleanupOldArchivedBatches() {
	archiveDir := filepath.Join(c.FailDir, "archived")
	_ = os.MkdirAll(archiveDir, 0755)
	cutoff := time.Now().AddDate(0, 0, -c.RetentionDays)
	filepath.WalkDir(archiveDir, func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() && filepath.Ext(path) == ".json" {
			if info, err := os.Stat(path); err == nil && info.ModTime().Before(cutoff) {
				log.WithField("file", path).Info("Deleting old archived batch")
				_ = os.Remove(path)
			}
		}
		return nil
	})
}

// RetryFailedBatches loads and resends failed batch files from the fail directory.
func (c *Client) RetryFailedBatches(ctx context.Context, queryParams map[string]string) error {
	c.cleanupOldArchivedBatches()

	entries, err := os.ReadDir(c.FailDir)
	if err != nil {
		return fmt.Errorf("cannot read fail dir: %w", err)
	}

	archiveDir := filepath.Join(c.FailDir, "archived")
	_ = os.MkdirAll(archiveDir, 0755)

	success := 0
	failures := 0

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		filePath := filepath.Join(c.FailDir, entry.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.WithError(err).WithField("file", filePath).Error("Could not read failed batch file")
			failures++
			continue
		}
		var payload tracker.NestedPayload
		if err := json.Unmarshal(data, &payload); err != nil {
			log.WithError(err).WithField("file", filePath).Error("Invalid JSON in failed batch file")
			failures++
			continue
		}

		log.WithField("file", filePath).Info("Retrying failed batch")
		_, _, err = c.SendTrackerPayload(ctx, &payload, queryParams)
		if err != nil {
			log.WithError(err).WithField("file", filePath).Error("Retry failed")
			failures++
			continue
		}

		archived := filepath.Join(archiveDir, entry.Name())
		if err := os.Rename(filePath, archived); err != nil {
			log.WithError(err).WithField("file", filePath).Error("Failed to archive retried batch file")
			failures++
			continue
		}
		log.WithField("file", archived).Info("Retry successful, archived file")
		success++
	}

	log.WithFields(log.Fields{
		"retried":  success + failures,
		"success":  success,
		"failures": failures,
	}).Info("Retry summary")

	if failures > 0 {
		return fmt.Errorf("%d retry failures occurred", failures)
	}
	return nil
}

// PollJobStatus retrieves the status of an async tracker import job by job ID.
func (c *Client) PollJobStatus(ctx context.Context, jobID string) (*schema.TrackerImportReport, *resty.Response, error) {
	if jobID == "" {
		return nil, nil, errors.New("job ID is empty")
	}

	req := c.Resty.R().SetContext(ctx).SetHeader("Accept", "application/json")
	res, err := req.Get(fmt.Sprintf("/api/tracker/jobs/%s/report", jobID))
	if err != nil {
		return nil, res, fmt.Errorf("failed to fetch job status: %w", err)
	}
	if res.StatusCode() != http.StatusOK {
		return nil, res, fmt.Errorf("unexpected status code: %d - %s", res.StatusCode(), res.String())
	}

	var report schema.TrackerImportReport
	if err := json.Unmarshal(res.Body(), &report); err != nil {
		return nil, res, fmt.Errorf("failed to decode job report: %w", err)
	}
	return &report, res, nil
}
