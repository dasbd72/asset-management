package bitfinexRest

import (
	"net/http"
)

type SecType int

const (
	// Version1 is for v1 API
	Version1 = "/v1"
	// Version2 is for v2 API
	Version2 = "/v2"
)

const (
	// SecTypePublic is for public API
	SecTypePublic SecType = iota
	// SecTypePrivate is for private API
	SecTypePrivate
)

// Request define an API request, build with Request_builder
type Request struct {
	method   string
	endpoint string
	version  string
	secType  SecType
	params   map[string]interface{}
}

// Request_builder define a builder for Request
type Request_builder struct {
	Method   string
	Endpoint string
	Version  string
	SecType  SecType
	Params   map[string]interface{}
}

// Build create a new Request
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
		version:  b.Version,
		secType:  b.SecType,
		params:   b.Params,
	}
}

// RequestOption define option type for request
type RequestOption func(*Request)
