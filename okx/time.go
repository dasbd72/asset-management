package okx

import "time"

func currentTimestamp() string {
	return formatTimestamp(time.Now().UTC())
}

// formatTimestamp formats a time into string.
func formatTimestamp(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.999Z07:00")
}
