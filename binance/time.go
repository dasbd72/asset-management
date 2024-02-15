package binance

import "time"

func currentTimestamp() int64 {
	return formatTimestamp(time.Now())
}

// formatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func formatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
