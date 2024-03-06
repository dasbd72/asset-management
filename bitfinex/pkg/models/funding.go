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

	FundingOffersWriteResponse struct {
		MTS          cast.NilOrInt    `json:"mts"`
		Type         cast.NilOrString `json:"type"`
		MessageID    cast.NilOrInt    `json:"messageID"`
		FundingOffer FundingOffer     `json:"funding_offer_array"`
		Code         cast.NilOrInt    `json:"code"`
		Status       cast.NilOrString `json:"status"`
		Text         cast.NilOrString `json:"text"`
	}

	FRRType string

	SubmitFundingOfferRequest struct {
		params map[string]interface{}
	}

	SubmitFundingOfferResponse = FundingOffersWriteResponse

	CancelFundingOfferRequest struct {
		params map[string]interface{}
	}

	CancelFundingOfferResponse = FundingOffersWriteResponse

	CancelAllFundingOfferRequest struct {
		params map[string]interface{}
	}

	CancelAllFundingOfferResponse struct {
		MTS    cast.NilOrInt    `json:"mts"`
		Type   cast.NilOrString `json:"type"`
		Status cast.NilOrString `json:"status"`
		Text   cast.NilOrString `json:"text"`
	}

	FundingInfo struct {
		Symbol       cast.NilOrString  `json:"symbol"`
		YieldLoan    cast.NilOrFloat64 `json:"yield_loan"`
		YieldLend    cast.NilOrFloat64 `json:"yield_lend"`
		DurationLoan cast.NilOrFloat64 `json:"duration_loan"`
		DurationLend cast.NilOrFloat64 `json:"duration_lend"`
	}
)

const (
	FRRLIMIT    FRRType = "LIMIT"       // Place an order at an explicit, static rate
	FRRDELTAFIX FRRType = "FRRDELTAFIX" // Place an order at the Flash Return Rate (FRR)
	FRRDELTAVAR FRRType = "FRRDELTAVAR" // Place an order at an implicit, dynamic rate, relative to the FRR
)

func (data *FundingStats) FromRaw(raw []byte) error {
	container := []interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *FundingStats) fromIf(v []interface{}) {
	for _, vv := range v {
		fundingStat := FundingStat{}
		fundingStat.fromIf(vv.([]interface{}))
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
	container := []interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *FundingOffers) fromIf(v []interface{}) {
	for _, vv := range v {
		fundingOffer := FundingOffer{}
		fundingOffer.fromIf(vv.([]interface{}))
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
	data.Flags = cast.IfToNilOrInt(v[9])
	data.Status = cast.IfToNilOrString(v[10])
	data.Rate = cast.IfToNilOrFloat64(v[14])
	data.Period = cast.IfToNilOrInt(v[15])
	data.Notify = cast.IfToNilOrInt(v[16])
	data.Hidden = cast.IfToNilOrInt(v[17])
	data.Renew = cast.IfToNilOrInt(v[19])
}

func NewSubmitFundingOfferRequest(frrType FRRType, symbol string, amount string, rate string, period int) *SubmitFundingOfferRequest {
	return &SubmitFundingOfferRequest{
		params: map[string]interface{}{
			"type":   frrType,
			"symbol": symbol,
			"amount": amount,
			"rate":   rate,
			"period": period,
		},
	}
}

func (data *SubmitFundingOfferRequest) Flag(flag int) *SubmitFundingOfferRequest {
	data.params["flags"] = flag
	return data
}

func (data *SubmitFundingOfferRequest) Params() map[string]interface{} {
	return data.params
}

func (data *FundingOffersWriteResponse) FromRaw(raw []byte) error {
	container := []interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *FundingOffersWriteResponse) fromIf(v []interface{}) {
	data.MTS = cast.IfToNilOrInt(v[0])
	data.Type = cast.IfToNilOrString(v[1])
	data.MessageID = cast.IfToNilOrInt(v[2])
	data.FundingOffer.fromIf(v[4].([]interface{}))
	data.Code = cast.IfToNilOrInt(v[5])
	data.Status = cast.IfToNilOrString(v[6])
	data.Text = cast.IfToNilOrString(v[7])
}

func NewCancelFundingOfferRequest(id int) *CancelFundingOfferRequest {
	return &CancelFundingOfferRequest{
		params: map[string]interface{}{
			"id": id,
		},
	}
}

func (data *CancelFundingOfferRequest) Params() map[string]interface{} {
	return data.params
}

func NewCancelAllFundingOfferRequest() *CancelAllFundingOfferRequest {
	return &CancelAllFundingOfferRequest{
		params: map[string]interface{}{},
	}
}

// Currency sets the currency for the request
//
// Example: "USD", "UST", "BTC"
func (data *CancelAllFundingOfferRequest) Currency(currency string) *CancelAllFundingOfferRequest {
	data.params["currency"] = currency
	return data
}

func (data *CancelAllFundingOfferRequest) Params() map[string]interface{} {
	return data.params
}

func (data *CancelAllFundingOfferResponse) FromRaw(raw []byte) error {
	container := []interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *CancelAllFundingOfferResponse) fromIf(v []interface{}) {
	data.MTS = cast.IfToNilOrInt(v[0])
	data.Type = cast.IfToNilOrString(v[1])
	data.Status = cast.IfToNilOrString(v[6])
	data.Text = cast.IfToNilOrString(v[7])
}

func (data *FundingInfo) FromRaw(raw []byte) error {
	container := []interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	data.fromIf(container)
	return nil
}

func (data *FundingInfo) fromIf(v []interface{}) {
	data.Symbol = cast.IfToNilOrString(v[1])
	data.YieldLoan = cast.IfToNilOrFloat64(v[2].([]interface{})[0])
	data.YieldLend = cast.IfToNilOrFloat64(v[2].([]interface{})[1])
	data.DurationLoan = cast.IfToNilOrFloat64(v[2].([]interface{})[2])
	data.DurationLend = cast.IfToNilOrFloat64(v[2].([]interface{})[3])
}
