package ws

type (
	SecType int
)

const (
	// SecTypeNone defines no security type
	SecTypeNone SecType = iota
	// SecTypeAPIKey defines signature by API key
	SecTypeAPIKey
	// SecTypeSigned defines signature by HMAC-SHA256
	SecTypeSigned
	// SecTypeEd25519 defines signature by ed25519
	SecTypeEd25519
)

type Request struct {
	method  string
	secType SecType
	params  map[string]interface{}
}

type Request_builder struct {
	Method  string
	SecType SecType
	Params  map[string]interface{}
}

// Build create a new Request
func (b Request_builder) Build() *Request {
	if b.Params == nil {
		b.Params = map[string]interface{}{}
	}
	return &Request{
		method:  b.Method,
		secType: b.SecType,
		params:  b.Params,
	}
}

// RequestOption define option type for request
type RequestOption func(*Request)
