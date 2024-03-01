package pionex

import (
	"fmt"
)

// APIError define API error when response status is 4xx or 5xx
type APIError struct {
	BaseResponse
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error return error code and message
func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> code=%s, message=%s", e.Code, e.Message)
}

// IsAPIError check if e is an API error
func IsAPIError(e error) bool {
	_, ok := e.(*APIError)
	return ok
}
