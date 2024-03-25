package models

import "github.com/dasbd72/go-exchange-sdk/okx/pkg/cast"

type BasicResponse struct {
	Code cast.JSONInt64 `json:"code"`
	Msg  string         `json:"msg"`
}
