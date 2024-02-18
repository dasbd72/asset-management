package max

import (
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestJSONFloat64(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    float64
		wantErr bool
	}{
		{
			name: "zero",
			in:   "0",
			want: 0,
		},
		{
			name: "negative",
			in:   "-1.0",
			want: -1.0,
		},
		{
			name: "empty",
			in:   "",
			want: 0,
		},
		{
			name: "float",
			in:   "0.10203101249012",
			want: 0.10203101249012,
		},
		{
			name:    "invalid",
			in:      "invalid",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var f JSONFloat64
			err := f.UnmarshalJSON([]byte(test.in))
			if test.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if f.Float64() != test.want {
				t.Errorf("expected %f, got %f", test.want, f.Float64())
			}
		})
	}
}

func TestJSONInt64(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    int64
		wantErr bool
	}{
		{
			name: "zero",
			in:   "0",
			want: 0,
		},
		{
			name: "negative",
			in:   "-1",
			want: -1,
		},
		{
			name: "empty",
			in:   "",
			want: 0,
		},
		{
			name: "int",
			in:   "123",
			want: 123,
		},
		{
			name:    "invalid",
			in:      "invalid",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var f JSONInt64
			err := f.UnmarshalJSON([]byte(test.in))
			if test.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if f.Int64() != test.want {
				t.Errorf("expected %d, got %d", test.want, f.Int64())
			}
		})
	}
}

func TestJSONTime(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    time.Time
		wantErr bool
	}{
		{
			name: "zero",
			in:   "0",
			want: time.UnixMilli(0),
		},
		{
			name: "one",
			in:   "1",
			want: time.UnixMilli(1),
		},
		{
			name: "empty",
			in:   "",
			want: time.Time{},
		},
		{
			name: "time",
			in:   "1694061154503",
			want: time.UnixMilli(1694061154503),
		},
		{
			name:    "invalid",
			in:      "invalid",
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var f JSONTime
			err := f.UnmarshalJSON([]byte(test.in))
			if test.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(f.Time(), test.want); diff != "" {
				t.Errorf("unexpected time (-got +want): %s", diff)
			}
			{
				got, err := f.MarshalJSON()
				if err != nil {
					t.Fatal(err)
				}
				if diff := cmp.Diff(string(got), strconv.FormatInt(test.want.UnixMilli(), 10)); diff != "" {
					t.Errorf("unexpected time (-got +want): %s", diff)
				}
			}
			if diff := cmp.Diff(f.String(), test.want.String()); diff != "" {
				t.Errorf("unexpected time (-got +want): %s", diff)
			}
		})
	}
}
