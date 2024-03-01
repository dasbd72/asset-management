package bitfinex

import (
	"fmt"
	"time"
)

func currentTimestamp() string {
	return fmt.Sprintf("%d", formatTimestamp(time.Now()))
}

// formatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func formatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
