package bitfinex

import (
	"encoding/json"
	"fmt"
)

// APIError define API error
//
// Refer to https://binance-docs.github.io/apidocs/spot/en/#error-codes
// for error code details
type APIError []interface{}

// Error return error code and message
func (e APIError) Error() string {
	if len(e) == 3 {
		return fmt.Sprintf("<APIError> code=%v, msg=\"%v\"", e[1], e[2])
	}
	b, _ := json.Marshal(e)
	return fmt.Sprintf("<APIError> msg=%s", string(b))
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
