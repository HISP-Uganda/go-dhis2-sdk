package dhis2

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestPollJobStatus_Success(t *testing.T) {
	client := NewClient("http://localhost:8080", "admin", "district")

	client.Resty = resty.New().SetTransport(mockRoundTripper(func(req *http.Request) (*http.Response, error) {
		if req.URL.Path == "/api/tracker/jobs/abc123/report" {
			return mockResponse(200, `{"status":"OK","importCount":{"imported":1}}`), nil
		}
		return nil, errors.New("unexpected path")
	}))

	report, resp, err := client.PollJobStatus(context.Background(), "abc123")
	assert.NoError(t, err)
	assert.NotNil(t, report)
	assert.NotNil(t, resp)
	// assert.Equal(t, 1, report.ImportCount.Imported)
}

func TestPollJobStatus_EmptyJobID(t *testing.T) {
	client := NewClient("http://localhost:8080", "admin", "district")
	report, resp, err := client.PollJobStatus(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, report)
	assert.Nil(t, resp)
}

func TestPollJobStatus_BadStatus(t *testing.T) {
	client := NewClient("http://localhost:8080", "admin", "district")

	client.Resty = resty.New().SetTransport(mockRoundTripper(func(req *http.Request) (*http.Response, error) {
		return mockResponse(500, `Internal Server Error`), nil
	}))

	report, resp, err := client.PollJobStatus(context.Background(), "abc123")
	assert.Error(t, err)
	assert.Nil(t, report)
	assert.NotNil(t, resp)
}

// mock helpers
func mockResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		Body:       ioNopCloser(body),
		Header:     make(http.Header),
	}
}

type mockRoundTripper func(req *http.Request) (*http.Response, error)

func (m mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m(req)
}

func ioNopCloser(s string) *nopCloser {
	return &nopCloser{data: []byte(s)}
}

type nopCloser struct {
	data []byte
	pos  int
}

func (n *nopCloser) Read(p []byte) (int, error) {
	if n.pos >= len(n.data) {
		return 0, io.EOF
	}
	nn := copy(p, n.data[n.pos:])
	n.pos += nn
	return nn, nil
}

func (n *nopCloser) Close() error {
	return nil
}
