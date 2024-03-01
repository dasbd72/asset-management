package pionex

import "net/http"

type SecType int

const (
	// SecTypePublic defines no security type
	SecTypePublic SecType = iota
	// SecTypePrivate defines signature required:
	//
	// timestamp(query) defines the current timestamp in milliseconds
	//
	// PIONEX-SIGNATURE(header) defines the signature of the request
	SecTypePrivate
)

// Request define an API request, build with Request_builder
type Request struct {
	method   string
	endpoint string
	secType  SecType
	params   map[string]interface{}
}

// Request_builder define a builder for Request
type Request_builder struct {
	Method   string
	Endpoint string
	SecType  SecType
	Params   map[string]interface{}
}

func (b Request_builder) Build() *Request {
	if b.Method == "" {
		b.Method = http.MethodGet
	}
	if b.Params == nil {
		b.Params = map[string]interface{}{}
	}
	return &Request{
		method:   b.Method,
		endpoint: b.Endpoint,
		secType:  b.SecType,
		params:   b.Params,
	}
}

// RequestOption define option type for request
type RequestOption func(*Request)
