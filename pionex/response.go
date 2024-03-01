package pionex

type BaseResponse struct {
	Result    bool  `json:"result"`    // Response result
	Timestamp int64 `json:"timestamp"` // Response timestamp in milliseconds
}
