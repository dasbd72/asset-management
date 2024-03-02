package bitfinexRest

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestError(t *testing.T) {
	tests := []struct {
		name string
		e    APIError
		want string
	}{
		{
			name: "empty",
			e:    APIError{},
			want: "<APIError> msg=[]",
		},
		{
			name: "non-empty",
			e:    APIError{"error", 10000, "error: test message"},
			want: "<APIError> code=10000, msg=\"error: test message\"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.e.Error()
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("Error() got unexpected result: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestIsAPIError(t *testing.T) {
	tests := []struct {
		name string
		e    error
		want bool
	}{
		{
			name: "nil",
			e:    nil,
			want: false,
		},
		{
			name: "APIError",
			e:    &APIError{},
			want: true,
		},
		{
			name: "other",
			e:    fmt.Errorf("test"),
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsAPIError(test.e)
			if got != test.want {
				t.Errorf("IsAPIError() = %v, want %v", got, test.want)
			}
		})
	}
}
