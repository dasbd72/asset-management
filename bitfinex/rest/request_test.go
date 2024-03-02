package bitfinexRest

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
				method:  http.MethodGet,
				secType: SecTypePublic,
				params:  map[string]interface{}{},
			},
		},
		{
			name: "non-empty",
			b: Request_builder{
				Method:   http.MethodPost,
				Endpoint: "/test",
				SecType:  SecTypePrivate,
				Params:   map[string]interface{}{"test": "test"},
			},
			want: &Request{
				method:   http.MethodPost,
				endpoint: "/test",
				secType:  SecTypePrivate,
				params:   map[string]interface{}{"test": "test"},
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
