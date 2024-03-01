package pionex

import (
	"testing"
	"time"
)

func Test_formatTimestamp(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1970-01-01 00:00:00 (epoch)",
			args: args{
				t: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "2016-06-01 01:01:01",
			args: args{
				t: time.Date(2016, 6, 1, 1, 1, 1, 0, time.UTC),
			},
			want: 1464742861000,
		},
		{
			name: "2018-06-01 01:01:01",
			args: args{
				t: time.Date(2018, 6, 1, 1, 1, 1, 0, time.UTC),
			},
			want: 1527814861000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTimestamp(tt.args.t); got != tt.want {
				t.Errorf("formatTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
