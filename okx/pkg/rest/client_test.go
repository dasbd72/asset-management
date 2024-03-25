package rest

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_sign(t *testing.T) {
	type args struct {
		secret  string
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testcase 1",
			args: args{
				secret:  "secret",
				message: "message",
			},
			want: "i19IcCmVwVmMVz2x4hhmqbgl1KeU0WnXBgoDYFeWNgs=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sign(tt.args.secret, tt.args.message); got != tt.want {
				t.Errorf("sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimestamp(t *testing.T) {
	tpe, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		t.Fatalf("time.LoadLocation() failed: %v", err)
	}
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "zero",
			tm:   time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			want: "1970-01-01T00:00:00Z",
		},
		{
			name: "one",
			tm:   time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC),
			want: "1970-01-01T00:00:01Z",
		},
		{
			name: "dev time",
			tm:   time.Date(2024, 2, 18, 22, 34, 0, 0, tpe),
			want: "2024-02-18T22:34:00+08:00",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := formatTimestamp(test.tm)
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("formatTimestamp() got unexpected result: (-got, +want)\n%s", diff)
			}
		})
	}
}
