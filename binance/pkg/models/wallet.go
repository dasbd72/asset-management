package models

import "github.com/dasbd72/go-exchange-sdk/binance/pkg/cast"

type (
	GetWalletStatusResponse struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	}

	GetUserWalletBalanceResponse []struct {
		Activate   bool             `json:"activate"`
		Balance    cast.JSONFloat64 `json:"balance"` // in BTC
		WalletName string           `json:"walletName"`
	}
)
