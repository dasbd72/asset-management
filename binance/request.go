package binance

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

type SecType int

const (
	// SecTypeNone defines no security type
	SecTypeNone SecType = iota
	// SecTypeAPIKey defines API key is required
	SecTypeAPIKey
	// SecTypeSigned defines both API key and signature is required
	SecTypeSigned // if the 'timestamp' parameter is required
)

type Params map[string]interface{}

// Request define an API request, build with Request_builder
type Request struct {
	method     string
	endpoint   string
	query      url.Values
	form       url.Values
	recvWindow int64
	secType    SecType
	header     http.Header
	body       io.Reader
	fullURL    string
}

// Request_builder define a builder for Request
type Request_builder struct {
	Method   string
	Endpoint string
	SecType  SecType
}

// Build create a new Request
func (b Request_builder) Build() *Request {
	return &Request{
		method:     b.Method,
		endpoint:   b.Endpoint,
		query:      url.Values{},
		form:       url.Values{},
		recvWindow: 0,
		secType:    b.SecType,
		header:     http.Header{},
		body:       nil,
		fullURL:    "",
	}
}

// AddParam add param with key/value to query string
func (r *Request) AddParam(key string, value interface{}) *Request {
	r.query.Add(key, fmt.Sprintf("%v", value))
	return r
}

// SetParam set param with key/value to query string
func (r *Request) SetParam(key string, value interface{}) *Request {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		v, err := json.Marshal(value)
		if err == nil {
			value = string(v)
		}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// SetParams set Params with key/values to query string
func (r *Request) SetParams(m Params) *Request {
	for k, v := range m {
		r.SetParam(k, v)
	}
	return r
}

// SetFormParam set param with key/value to request form body
func (r *Request) SetFormParam(key string, value interface{}) *Request {
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

// SetFormParams set Params with key/values to request form body
func (r *Request) SetFormParams(m Params) *Request {
	for k, v := range m {
		r.SetFormParam(k, v)
	}
	return r
}

// RequestOption define option type for request
type RequestOption func(*Request)

// WithRecvWindow set recvWindow param for the request
func WithRecvWindow(recvWindow int64) RequestOption {
	return func(r *Request) {
		r.recvWindow = recvWindow
	}
}

// WithHeader set or add a header value to the request
func WithHeader(key, value string, replace bool) RequestOption {
	return func(r *Request) {
		if replace {
			r.header.Set(key, value)
		} else {
			r.header.Add(key, value)
		}
	}
}

// WithHeaders set or replace the headers of the request
func WithHeaders(header http.Header) RequestOption {
	return func(r *Request) {
		r.header = header.Clone()
	}
}
