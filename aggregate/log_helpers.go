package aggregate

import (
	log "github.com/sirupsen/logrus"
)

// LogSummary logs the status, import count, and any conflicts using logrus.
func (r *ImportSummaryResponse) LogSummary() {
	fields := log.Fields{
		"status":        r.Status,
		"imported":      r.ImportCount.Imported,
		"updated":       r.ImportCount.Updated,
		"ignored":       r.ImportCount.Ignored,
		"deleted":       r.ImportCount.Deleted,
		"conflictCount": len(r.Conflicts),
	}

	switch r.Status {
	case "SUCCESS":
		if len(r.Conflicts) > 0 {
			log.WithFields(fields).Warn("DHIS2 submission succeeded but had warnings.")
			r.LogConflicts()
		} else {
			log.WithFields(fields).Info("DHIS2 submission succeeded.")
		}
	case "WARNING":
		log.WithFields(fields).Warn("DHIS2 submission completed with warnings.")
		r.LogConflicts()
	case "ERROR":
		log.WithFields(fields).Error("DHIS2 submission failed with errors.")
		r.LogConflicts()
	default:
		log.WithFields(fields).Info("DHIS2 submission returned unknown status.")
	}
}

// LogConflicts logs individual conflicts in detail.
func (r *ImportSummaryResponse) LogConflicts() {
	for _, c := range r.Conflicts {
		log.WithFields(log.Fields{
			"property":   c.Property,
			"value":      c.Value,
			"errorCode":  c.ErrorCode,
			"indexes":    c.Indexes,
			"objectHint": c.Objects,
		}).Warn("Conflict in DHIS2 submission")
	}
}
