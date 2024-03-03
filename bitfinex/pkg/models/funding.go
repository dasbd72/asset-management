package models

import (
	"encoding/json"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/cast"
)

type (
	FundingStats struct {
		FundingStats []FundingStat `json:"funding_stat_array"`
	}

	FundingStat struct {
		MTS                         cast.NilOrInt     `json:"mts"`
		FRR                         cast.NilOrFloat64 `json:"frr"`
		AvgPeriod                   cast.NilOrFloat64 `json:"avg_period"`
		FundingAmount               cast.NilOrFloat64 `json:"funding_amount"`
		FundingAmountUsed           cast.NilOrFloat64 `json:"funding_amount_used"`
		FundingAmountBelowThreshold cast.NilOrFloat64 `json:"funding_amount_below_threshold"`
	}

	FundingOffers struct {
		FundingOffers []FundingOffer `json:"funding_offer_array"`
	}

	FundingOffer struct {
		ID         cast.NilOrInt     `json:"id"`
		Symbol     cast.NilOrString  `json:"symbol"`
		MTSCreated cast.NilOrInt     `json:"mts_create"`
		MTSUpdated cast.NilOrInt     `json:"mts_update"`
		Amount     cast.NilOrFloat64 `json:"amount"`
		AmountOrig cast.NilOrFloat64 `json:"amount_orig"`
		Type       cast.NilOrString  `json:"type"`
		Flags      interface{}       `json:"flags"`
		Status     cast.NilOrString  `json:"status"`
		Rate       cast.NilOrFloat64 `json:"rate"`
		Period     cast.NilOrInt     `json:"period"`
		Notify     cast.NilOrInt     `json:"notify"`
		Hidden     cast.NilOrInt     `json:"hidden"`
		Renew      cast.NilOrInt     `json:"renew"`
	}
)

func (data *FundingStats) FromRaw(raw []byte) error {
	container := [][]interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *FundingStats) fromIf(v [][]interface{}) {
	for _, vv := range v {
		fundingStat := FundingStat{}
		fundingStat.fromIf(vv)
	}
}

func (data *FundingStat) FromRaw(raw []byte) error {
	container := []interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *FundingStat) fromIf(v []interface{}) {
	data.MTS = cast.IfToNilOrInt(v[0])
	data.FRR = cast.IfToNilOrFloat64(v[3])
	data.AvgPeriod = cast.IfToNilOrFloat64(v[4])
	data.FundingAmount = cast.IfToNilOrFloat64(v[7])
	data.FundingAmountUsed = cast.IfToNilOrFloat64(v[8])
	data.FundingAmountBelowThreshold = cast.IfToNilOrFloat64(v[11])
}

func (data *FundingOffers) FromRaw(raw []byte) error {
	container := [][]interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *FundingOffers) fromIf(v [][]interface{}) {
	for _, vv := range v {
		fundingOffer := FundingOffer{}
		fundingOffer.fromIf(vv)
		data.FundingOffers = append(data.FundingOffers, fundingOffer)
	}
}

func (data *FundingOffer) FromRaw(raw []byte) error {
	container := []interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *FundingOffer) fromIf(v []interface{}) {
	data.ID = cast.IfToNilOrInt(v[0])
	data.Symbol = cast.IfToNilOrString(v[1])
	data.MTSCreated = cast.IfToNilOrInt(v[2])
	data.MTSUpdated = cast.IfToNilOrInt(v[3])
	data.Amount = cast.IfToNilOrFloat64(v[4])
	data.AmountOrig = cast.IfToNilOrFloat64(v[5])
	data.Type = cast.IfToNilOrString(v[6])
	data.Flags = v[7]
	data.Status = cast.IfToNilOrString(v[8])
	data.Rate = cast.IfToNilOrFloat64(v[9])
	data.Period = cast.IfToNilOrInt(v[10])
	data.Notify = cast.IfToNilOrInt(v[11])
	data.Hidden = cast.IfToNilOrInt(v[12])
	data.Renew = cast.IfToNilOrInt(v[13])
}
