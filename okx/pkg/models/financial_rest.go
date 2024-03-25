package models

import "github.com/dasbd72/go-exchange-sdk/okx/pkg/cast"

type (
	GetEarnOffersRequest struct {
		params map[string]interface{}
	}

	GetEarnOffersResponse struct {
		BasicResponse
		Offers []struct {
			Ccy          string           `json:"ccy"`
			ProductId    cast.JSONInt64   `json:"productId"`
			Protocol     string           `json:"protocol"`
			ProtocolType string           `json:"protocolType"`
			Term         cast.JSONInt64   `json:"term"`
			Apy          cast.JSONFloat64 `json:"apy"`
			EarlyRedeem  bool             `json:"earlyRedeem"`
			InvestData   []struct {
				Ccy    string           `json:"ccy"`
				Bal    cast.JSONFloat64 `json:"bal"`
				MinAmt cast.JSONFloat64 `json:"minAmt"`
				MaxAmt cast.JSONFloat64 `json:"maxAmt"`
			} `json:"investData"`
			EarningData []struct {
				Ccy         string         `json:"ccy"`
				EarningType cast.JSONInt64 `json:"earningType"`
			} `json:"earningData"`
			State string `json:"state"`
		} `json:"data,omitempty"`
	}

	GetActiveEarnOrdersRequest struct {
		params map[string]interface{}
	}

	GetActiveEarnOrdersResponse struct {
		BasicResponse
		Orders []struct {
			Ccy          string           `json:"ccy"`
			OrdId        cast.JSONInt64   `json:"ordId"`
			ProductId    cast.JSONInt64   `json:"productId"`
			State        string           `json:"state"`
			Protocol     string           `json:"protocol"`
			ProtocolType string           `json:"protocolType"`
			Term         cast.JSONInt64   `json:"term"`
			Apy          cast.JSONFloat64 `json:"apy"`
			InvestData   []struct {
				Ccy string           `json:"ccy"`
				Amt cast.JSONFloat64 `json:"amt"`
			} `json:"investData"`
			EarningData []struct {
				Ccy         string           `json:"ccy"`
				EarningType cast.JSONInt64   `json:"earningType"`
				Earnings    cast.JSONFloat64 `json:"earnings"`
			} `json:"earningData"`
			PurchasedTime            cast.JSONTime `json:"purchasedTime"`
			EstSettlementTime        cast.JSONTime `json:"estSettlementTime,omitempty"`
			CancelRedemptionDeadline cast.JSONTime `json:"cancelRedemptionDeadline,omitempty"`
			Tag                      string        `json:"tag,omitempty"`
		} `json:"data,omitempty"`
	}

	GetETHStakingBalanceResponse struct {
		BasicResponse
		Balances []struct {
			Ccy                   string           `json:"ccy"`
			Amt                   cast.JSONFloat64 `json:"amt"`
			LatestInterestAccrual cast.JSONFloat64 `json:"latestInterestAccrual"`
			TotalInterestAccrual  cast.JSONFloat64 `json:"totalInterestAccrual"`
			Ts                    cast.JSONTime    `json:"ts"`
		} `json:"data,omitempty"`
	}

	GetSavingBalanceRequest struct {
		params map[string]interface{}
	}

	GetSavingBalanceResponse struct {
		BasicResponse
		Balances []struct {
			Ccy        string           `json:"ccy"`
			Amt        cast.JSONFloat64 `json:"amt"`
			Earnings   cast.JSONFloat64 `json:"earnings"`
			Rate       cast.JSONFloat64 `json:"rate"`
			LoanAmt    cast.JSONFloat64 `json:"loanAmt"`
			PendingAmt cast.JSONFloat64 `json:"pendingAmt"`
			RedemptAmt cast.JSONFloat64 `json:"redemptAmt,omitempty"` // Deprecated
		} `json:"data,omitempty"`
	}

	GetLendingHistoryRequest struct {
		params map[string]interface{}
	}

	GetLendingHistoryResponse struct {
		BasicResponse
		Records []struct {
			Ccy      string           `json:"ccy"`
			Amt      cast.JSONFloat64 `json:"amt"`
			Earnings cast.JSONFloat64 `json:"earnings"`
			Rate     cast.JSONFloat64 `json:"rate"`
			Ts       cast.JSONTime    `json:"ts"`
		} `json:"data,omitempty"`
	}
)

func NewGetEarnOffersRequest() *GetEarnOffersRequest {
	return &GetEarnOffersRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetEarnOffersRequest) ProductId(productId string) *GetEarnOffersRequest {
	r.params["productId"] = productId
	return r
}

func (r *GetEarnOffersRequest) ProtocolType(protocolType string) *GetEarnOffersRequest {
	r.params["protocolType"] = protocolType
	return r
}

func (r *GetEarnOffersRequest) Ccy(ccy string) *GetEarnOffersRequest {
	r.params["ccy"] = ccy
	return r
}

func (r *GetEarnOffersRequest) Params() map[string]interface{} {
	return r.params
}

func NewGetActiveEarnOrdersRequest() *GetActiveEarnOrdersRequest {
	return &GetActiveEarnOrdersRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetActiveEarnOrdersRequest) ProductId(productId string) *GetActiveEarnOrdersRequest {
	r.params["productId"] = productId
	return r
}

func (r *GetActiveEarnOrdersRequest) ProtocolType(protocolType string) *GetActiveEarnOrdersRequest {
	r.params["protocolType"] = protocolType
	return r
}

func (r *GetActiveEarnOrdersRequest) Ccy(ccy string) *GetActiveEarnOrdersRequest {
	r.params["ccy"] = ccy
	return r
}

func (r *GetActiveEarnOrdersRequest) State(state cast.JSONInt64) *GetActiveEarnOrdersRequest {
	r.params["state"] = state
	return r
}

func (r *GetActiveEarnOrdersRequest) Params() map[string]interface{} {
	return r.params
}

func NewGetSavingBalanceRequest() *GetSavingBalanceRequest {
	return &GetSavingBalanceRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetSavingBalanceRequest) Ccy(ccy string) *GetSavingBalanceRequest {
	r.params["ccy"] = ccy
	return r
}

func (r *GetSavingBalanceRequest) Params() map[string]interface{} {
	return r.params
}

func NewGetLendingHistoryRequest() *GetLendingHistoryRequest {
	return &GetLendingHistoryRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetLendingHistoryRequest) Ccy(ccy string) *GetLendingHistoryRequest {
	r.params["ccy"] = ccy
	return r
}

// After sets pagination of data to return records earlier than the requested ts,
// Unix timestamp format in milliseconds, e.g. 1597026383085
func (r *GetLendingHistoryRequest) After(after int64) *GetLendingHistoryRequest {
	r.params["after"] = after
	return r
}

// Before sets pagination of data to return records newer than the requested ts,
// Unix timestamp format in milliseconds, e.g. 1597026383085
func (r *GetLendingHistoryRequest) Before(before int64) *GetLendingHistoryRequest {
	r.params["before"] = before
	return r
}

// Limit sets number of results per request. The maximum is 100. The default is 100.
func (r *GetLendingHistoryRequest) Limit(limit int64) *GetLendingHistoryRequest {
	r.params["limit"] = limit
	return r
}

func (r *GetLendingHistoryRequest) Params() map[string]interface{} {
	return r.params
}
