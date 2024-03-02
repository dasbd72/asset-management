package bitfinexRest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/cast"
)

type (
	GetFundingStats struct {
		FundingStatArray []FundingStat `json:"funding_stat_array"`
	}
	FundingStat struct {
		MTS                         cast.NilOrInt     `json:"mts"`
		FRR                         cast.NilOrFloat64 `json:"frr"`
		AvgPeriod                   cast.NilOrFloat64 `json:"avg_period"`
		FundingAmount               cast.NilOrFloat64 `json:"funding_amount"`
		FundingAmountUsed           cast.NilOrFloat64 `json:"funding_amount_used"`
		FundingAmountBelowThreshold cast.NilOrFloat64 `json:"funding_amount_below_threshold"`
	}

	GetActiveFundingOffers struct {
		FundingOfferArray []FundingOffer `json:"funding_offer_array"`
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

func (data *GetFundingStats) FromRaw(raw []byte) error {
	container := [][]interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	for _, v := range container {
		for i, vv := range v {
			if vv == nil {
				v[i] = ""
			}
		}
		data.FundingStatArray = append(data.FundingStatArray, FundingStat{
			MTS:                         cast.IfToNilOrInt(v[0]),
			FRR:                         cast.IfToNilOrFloat64(v[3]),
			AvgPeriod:                   cast.IfToNilOrFloat64(v[4]),
			FundingAmount:               cast.IfToNilOrFloat64(v[7]),
			FundingAmountUsed:           cast.IfToNilOrFloat64(v[8]),
			FundingAmountBelowThreshold: cast.IfToNilOrFloat64(v[11]),
		})
	}
	return nil
}

func (c *Client) GetFundingStats(ctx context.Context, symbol string, opts ...RequestOption) (*GetFundingStats, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("/funding/stats/%s/hist", symbol),
		Version:  Version2,
		SecType:  SecTypePublic,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetFundingStats{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (data *GetActiveFundingOffers) FromRaw(raw []byte) error {
	container := [][]interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	for _, v := range container {
		for i, vv := range v {
			if vv == nil {
				v[i] = ""
			}
		}
		data.FundingOfferArray = append(data.FundingOfferArray, FundingOffer{
			ID:         cast.IfToNilOrInt(v[0]),
			Symbol:     cast.IfToNilOrString(v[1]),
			MTSCreated: cast.IfToNilOrInt(v[2]),
			MTSUpdated: cast.IfToNilOrInt(v[3]),
			Amount:     cast.IfToNilOrFloat64(v[4]),
			AmountOrig: cast.IfToNilOrFloat64(v[5]),
			Type:       cast.IfToNilOrString(v[6]),
			Flags:      cast.IfToNilOrString(v[7]),
			Status:     cast.IfToNilOrString(v[8]),
			Rate:       cast.IfToNilOrFloat64(v[9]),
			Period:     cast.IfToNilOrInt(v[10]),
			Notify:     cast.IfToNilOrInt(v[11]),
			Hidden:     cast.IfToNilOrInt(v[12]),
			Renew:      cast.IfToNilOrInt(v[13]),
		})
	}
	return nil
}

func (c *Client) GetActiveFundingOffers(ctx context.Context, symbol string, opts ...RequestOption) (*GetActiveFundingOffers, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: fmt.Sprintf("/auth/r/funding/offers/%s", symbol),
		Version:  Version2,
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetActiveFundingOffers{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
