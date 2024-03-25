package models

import "github.com/dasbd72/go-exchange-sdk/okx/pkg/cast"

type (
	GetBalanceRequest struct {
		params map[string]interface{}
	}

	GetBalanceResponse struct {
		BasicResponse
		Balances []*struct {
			UTime       cast.JSONTime    `json:"uTime"`
			TotalEq     cast.JSONFloat64 `json:"totalEq"`
			IsoEq       cast.JSONFloat64 `json:"isoEq,omitempty"`
			AdjEq       cast.JSONFloat64 `json:"adjEq,omitempty"`
			OrdFroz     cast.JSONFloat64 `json:"ordFroz,omitempty"`
			Imr         cast.JSONFloat64 `json:"imr,omitempty"`
			Mmr         cast.JSONFloat64 `json:"mmr,omitempty"`
			BorrowFroz  cast.JSONFloat64 `json:"borrowFroz,omitempty"`
			MgnRatio    cast.JSONFloat64 `json:"mgnRatio,omitempty"`
			NotionalUsd cast.JSONFloat64 `json:"notionalUsd,omitempty"`
			Upl         cast.JSONFloat64 `json:"upl,omitempty"`
			Details     []*struct {
				Ccy           string           `json:"ccy"`
				Eq            cast.JSONFloat64 `json:"eq"`
				CashBal       cast.JSONFloat64 `json:"cashBal"`
				UTime         cast.JSONTime    `json:"uTime"`
				IsoEq         cast.JSONFloat64 `json:"isoEq,omitempty"`
				AvailEq       cast.JSONFloat64 `json:"availEq,omitempty"`
				DisEq         cast.JSONFloat64 `json:"disEq"`
				FixedBal      cast.JSONFloat64 `json:"fixedBal,omitempty"`
				AvailBal      cast.JSONFloat64 `json:"availBal"`
				FrozenBal     cast.JSONFloat64 `json:"frozenBal"`
				OrdFrozen     cast.JSONFloat64 `json:"ordFrozen,omitempty"`
				Liab          cast.JSONFloat64 `json:"liab,omitempty"`
				Upl           cast.JSONFloat64 `json:"upl,omitempty"`
				UplLiab       cast.JSONFloat64 `json:"uplLiab,omitempty"`
				CrossLiab     cast.JSONFloat64 `json:"crossLiab,omitempty"`
				IsoLiab       cast.JSONFloat64 `json:"isoLiab,omitempty"`
				MgnRatio      cast.JSONFloat64 `json:"mgnRatio,omitempty"`
				Interest      cast.JSONFloat64 `json:"interest,omitempty"`
				Twap          cast.JSONFloat64 `json:"twap,omitempty"`
				MaxLoan       cast.JSONFloat64 `json:"maxLoan,omitempty"`
				EqUsd         cast.JSONFloat64 `json:"eqUsd"`
				BorrowFroz    cast.JSONFloat64 `json:"borrowFroz,omitempty"`
				NotionalLever cast.JSONFloat64 `json:"notionalLever,omitempty"`
				StgyEq        cast.JSONFloat64 `json:"stgyEq"`
				IsoUpl        cast.JSONFloat64 `json:"isoUpl,omitempty"`
				SpotInUseAmt  cast.JSONFloat64 `json:"spotInUseAmt,omitempty"`
				SpotIsoBal    cast.JSONFloat64 `json:"spotIsoBal,omitempty"`
				Imr           cast.JSONFloat64 `json:"imr,omitempty"`
				Mmr           cast.JSONFloat64 `json:"mmr,omitempty"`
			} `json:"details,omitempty"`
		} `json:"data,omitempty"`
	}
)

func NewGetBalanceRequest() *GetBalanceRequest {
	return &GetBalanceRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetBalanceRequest) Ccy(ccy string) *GetBalanceRequest {
	r.params["ccy"] = ccy
	return r
}

func (r *GetBalanceRequest) Params() map[string]interface{} {
	return r.params
}
