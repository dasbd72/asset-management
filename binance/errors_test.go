package binance

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
			want: "<APIError> code=0, msg=",
		},
		{
			name: "non-empty",
			e:    APIError{Code: 1, Message: "test"},
			want: "<APIError> code=1, msg=test",
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
