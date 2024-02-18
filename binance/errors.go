package binance

import (
	"fmt"
)

// APIError define API error
//
// Refer to https://binance-docs.github.io/apidocs/spot/en/#error-codes
// for error code details
type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

// Error return error code and message
func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> code=%d, msg=%s", e.Code, e.Message)
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
