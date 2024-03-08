package models

import (
	"encoding/json"
	"fmt"
)

type (
	GetCandlesRequest struct {
		Candle  string
		Section CandleSectionType
		params  map[string]interface{}
	}
	Candles []Candle
	Candle  []interface{}
)

type CandleTimeFrame string

const (
	CandleTFOneMinute      CandleTimeFrame = "1m"
	CandleTFFiveMinutes    CandleTimeFrame = "5m"
	CandleTFFifteenMinutes CandleTimeFrame = "15m"
	CandleTFThirtyMinutes  CandleTimeFrame = "30m"
	CandleTFOneHour        CandleTimeFrame = "1h"
	CandleTFThreeHours     CandleTimeFrame = "3h"
	CandleTFSixHours       CandleTimeFrame = "6h"
	CandleTFTwelveHours    CandleTimeFrame = "12h"
	CandleTFOneDay         CandleTimeFrame = "1D"
	CandleTFOneWeek        CandleTimeFrame = "7D"
	CandleTFTwoWeeks       CandleTimeFrame = "14D"
	CandleTFOneMonth       CandleTimeFrame = "1M"
)

type CandleSectionType string

const (
	CandleSTHist CandleSectionType = "hist"
	CandleSTLast CandleSectionType = "last"
)

// NewGetCandlesRequest creates a new instance of GetCandlesRequest
//
// Refer to https://docs.bitfinex.com/reference/rest-public-candles
// for available trading pairs and time frames
func NewGetCandlesRequest(candle string, section CandleSectionType) *GetCandlesRequest {
	return &GetCandlesRequest{
		Candle:  candle,
		Section: section,
		params:  make(map[string]interface{}),
	}
}

// NewGetTradingPairsCandlesRequest creates a new instance of GetCandlesRequest for trading pairs
func NewGetTradingPairsCandlesRequest(timeFrame CandleTimeFrame, symbol string, section CandleSectionType) *GetCandlesRequest {
	return NewGetCandlesRequest(fmt.Sprintf("trade:%s:%s", timeFrame, symbol), section)
}

// NewGetFundingCurrenciesCandlesRequest creates a new instance of GetCandlesRequest for funding currencies
//
// period: 2 to 120 (days)
func NewGetFundingCurrenciesCandlesRequest(timeFrame CandleTimeFrame, symbol string, period int, section CandleSectionType) *GetCandlesRequest {
	return NewGetCandlesRequest(fmt.Sprintf("trade:%s:%s:p%d", timeFrame, symbol, period), section)
}

// Sort sets sorting order
//
// +1: sort in ascending order | -1: sort in descending order (by MTS field).
func (data *GetCandlesRequest) Sort(s int) *GetCandlesRequest {
	data.params["sort"] = s
	return data
}

// Start sets the filter for the start of the time range
//
// If start is given, only records with MTS >= start (milliseconds) will be given as response.
func (data *GetCandlesRequest) Start(s string) *GetCandlesRequest {
	data.params["start"] = s
	return data
}

// End sets the filter for the end of the time range
//
// If end is given, only records with MTS <= end (milliseconds) will be given as response.
func (data *GetCandlesRequest) End(e string) *GetCandlesRequest {
	data.params["end"] = e
	return data
}

// Limit sets the limit of candles to return
//
// Number of records in response (max. 10000).
func (data *GetCandlesRequest) Limit(l int) *GetCandlesRequest {
	data.params["limit"] = l
	return data
}

func (data *GetCandlesRequest) Params() map[string]interface{} {
	return data.params
}

func (data *Candles) FromRaw(raw []byte) error {
	err := json.Unmarshal(raw, data)
	if err != nil {
		return err
	}
	return nil
}

func (data Candle) GetMTS() int64 {
	return int64(data[0].(float64))
}

func (data Candle) GetOpen() float64 {
	return data[1].(float64)
}

func (data Candle) GetClose() float64 {
	return data[2].(float64)
}

func (data Candle) GetHigh() float64 {
	return data[3].(float64)
}

func (data Candle) GetLow() float64 {
	return data[4].(float64)
}

func (data Candle) GetVolume() float64 {
	return data[5].(float64)
}
