package binance

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name string
		b    Request_builder
		want *Request
	}{
		{
			name: "empty",
			b:    Request_builder{},
			want: &Request{
				method:     http.MethodGet,
				endpoint:   "",
				apiType:    ApiTypeSpot,
				secType:    SecTypeNone,
				recvWindow: 0,
				params:     map[string]interface{}{},
			},
		},
		{
			name: "non-empty",
			b: Request_builder{
				Method:   http.MethodPost,
				Endpoint: "/test",
				ApiType:  ApiTypeFutures,
				SecType:  SecTypeAPIKey,
				Params:   map[string]interface{}{"test": "test"},
			},
			want: &Request{
				method:     http.MethodPost,
				endpoint:   "/test",
				apiType:    ApiTypeFutures,
				secType:    SecTypeAPIKey,
				recvWindow: 0,
				params:     map[string]interface{}{"test": "test"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.b.Build()
			if diff := cmp.Diff(got, test.want, cmp.AllowUnexported(Request{})); diff != "" {
				t.Errorf("Build() got unexpected result: (-got, +want)\n%s", diff)
			}
		})
	}
}

func TestWithRecvWindow(t *testing.T) {
	r := &Request{}
	WithRecvWindow(1000)(r)
	if r.recvWindow != 1000 {
		t.Errorf("WithRecvWindow() recvWindow got unexpected result: got=%d, want=%d", r.recvWindow, 1000)
	}
}
