package models

import (
	"encoding/json"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/cast"
)

type (
	Wallets struct {
		Wallets []Wallet `json:"wallet"`
	}

	Wallet struct {
		Type               cast.NilOrString  `json:"type"`
		Currency           cast.NilOrString  `json:"currency"`
		Balance            cast.NilOrFloat64 `json:"balance"`
		UnsettledInterest  cast.NilOrFloat64 `json:"unsettled_interest"`
		AvailableBalance   cast.NilOrFloat64 `json:"available_balance"`
		LastChange         cast.NilOrString  `json:"last_change"`
		LastChangeMetadata interface{}       `json:"last_change_metadata"`
	}
)

func (data *Wallets) FromRaw(raw []byte) error {
	container := [][]interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *Wallets) fromIf(v [][]interface{}) {
	for _, vv := range v {
		wallet := Wallet{}
		wallet.fromIf(vv)
		data.Wallets = append(data.Wallets, wallet)
	}
}

func (data *Wallet) FromRaw(raw []byte) error {
	container := []interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *Wallet) fromIf(v []interface{}) {
	data.Type = cast.IfToNilOrString(v[0])
	data.Currency = cast.IfToNilOrString(v[1])
	data.Balance = cast.IfToNilOrFloat64(v[2])
	data.UnsettledInterest = cast.IfToNilOrFloat64(v[3])
	data.AvailableBalance = cast.IfToNilOrFloat64(v[4])
	data.LastChange = cast.IfToNilOrString(v[5])
	data.LastChangeMetadata = v[6]
}
