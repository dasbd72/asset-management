package pionex

import (
	"fmt"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	type fields struct {
		BaseResponse BaseResponse
		Code         string
		Message      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty",
			want: "<APIError> code=, message=",
		},
		{
			name: "non-empty",
			fields: fields{
				Code:    "APIKEY_LOST",
				Message: "test",
			},
			want: "<APIError> code=APIKEY_LOST, message=test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := APIError{
				BaseResponse: tt.fields.BaseResponse,
				Code:         tt.fields.Code,
				Message:      tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("APIError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAPIError(t *testing.T) {
	type args struct {
		e error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{
				e: nil,
			},
			want: false,
		},
		{
			name: "APIError",
			args: args{
				e: &APIError{},
			},
			want: true,
		},
		{
			name: "other",
			args: args{
				fmt.Errorf("test"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAPIError(tt.args.e); got != tt.want {
				t.Errorf("IsAPIError() = %v, want %v", got, tt.want)
			}
		})
	}
}
