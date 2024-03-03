package spot

type (
	SecType int
	ApiType int
)

const (
	// SecTypeNone defines no security type
	SecTypeNone SecType = iota
	// SecTypeAPIKey defines API key is required
	SecTypeAPIKey
	// SecTypeSigned defines both API key and signature is required
	SecTypeSigned // if the 'timestamp' parameter is required
)

const (
	// ApiTypeSpot uses baseAPIMainURL
	ApiTypeSpot ApiType = iota
	// ApiTypeFutures uses futuresAPIMainURL
	ApiTypeFutures
)

// Request define an API request, build with Request_builder
type Request struct {
	method     string
	endpoint   string
	apiType    ApiType
	secType    SecType
	recvWindow int64
	params     map[string]interface{}
}

// Request_builder define a builder for Request
type Request_builder struct {
	Method   string
	Endpoint string
	ApiType  ApiType
	SecType  SecType
	Params   map[string]interface{}
}

// Build create a new Request
func (b Request_builder) Build() *Request {
	if b.Method == "" {
		b.Method = "GET"
	}
	if b.Params == nil {
		b.Params = map[string]interface{}{}
	}
	return &Request{
		method:     b.Method,
		endpoint:   b.Endpoint,
		apiType:    b.ApiType,
		secType:    b.SecType,
		recvWindow: 0,
		params:     b.Params,
	}
}

// RequestOption define option type for request
type RequestOption func(*Request)

// WithRecvWindow set recvWindow param for the request
func WithRecvWindow(recvWindow int64) RequestOption {
	return func(r *Request) {
		r.recvWindow = recvWindow
	}
}
