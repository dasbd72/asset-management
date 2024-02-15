package binance

import (
	"testing"
	"time"
)

func TestFormatTimestamp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		t    time.Time
		want int64
	}{
		{
			name: "1970-01-01 00:00:00 (epoch)",
			t:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			want: 0,
		},
		{
			name: "2016-06-01 01:01:01",
			t:    time.Date(2016, 6, 1, 1, 1, 1, 0, time.UTC),
			want: 1464742861000,
		},
		{
			name: "2018-06-01 01:01:01",
			t:    time.Date(2018, 6, 1, 1, 1, 1, 0, time.UTC),
			want: 1527814861000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := formatTimestamp(test.t)
			if got != test.want {
				t.Errorf("formatTimestamp(%v) got %v, want %v", test.t, got, test.want)
			}
		})
	}
}
