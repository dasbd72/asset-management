package okx

type SecType int

const (
	// SecTypePublic defines no security type
	SecTypePublic SecType = iota
	// SecTypePrivate defines headers and signature required:
	//
	// OK-ACCESS-KEY The API Key as a String.
	//
	// OK-ACCESS-SIGN The Base64-encoded signature (see Signing Messages subsection for details).
	//
	// OK-ACCESS-TIMESTAMP The UTC timestamp of your request .e.g : 2020-12-08T09:08:57.715Z
	//
	// OK-ACCESS-PASSPHRASE The passphrase you specified when creating the APIKey.
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
	return &Request{
		method:   b.Method,
		endpoint: b.Endpoint,
		secType:  b.SecType,
		params:   b.Params,
	}
}

// SetParam set param with key/value to query string
func (r *Request) SetParam(key string, value interface{}) *Request {
	r.params[key] = value
	return r
}

// RequestOption define option type for request
type RequestOption func(*Request)
