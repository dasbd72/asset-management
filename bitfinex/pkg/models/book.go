package models

import "encoding/json"

type (
	GetBooksRequest struct {
		Symbol    string
		Precision BookPrecision
		params    map[string]interface{}
	}

	TradingPairBooks []TradingPairBook
	TradingPairBook  []interface{}

	FundingCurrencyBooks []FundingCurrencyBook
	FundingCurrencyBook  []interface{}
)

type BookPrecision string

const (
	BookPrecisionP0 BookPrecision = "P0"
	BookPrecisionP1 BookPrecision = "P1"
	BookPrecisionP2 BookPrecision = "P2"
	BookPrecisionP3 BookPrecision = "P3"
	BookPrecisionP4 BookPrecision = "P4"
	BookPrecisionR0 BookPrecision = "R0"
)

func NewGetBooksRequest(symbol string, precision BookPrecision) *GetBooksRequest {
	return &GetBooksRequest{
		Symbol:    symbol,
		Precision: precision,
		params:    make(map[string]interface{}),
	}
}

// Len sets the number of price points ("len") to return
//
// Must be 1, 25, 100
func (data *GetBooksRequest) Len(l int) *GetBooksRequest {
	data.params["len"] = l
	return data
}

func (data *GetBooksRequest) Params() map[string]interface{} {
	return data.params
}

func (data *TradingPairBooks) FromRaw(raw []byte) error {
	err := json.Unmarshal(raw, data)
	if err != nil {
		return err
	}
	return nil
}

func (data TradingPairBook) GetPrice() float64 {
	return data[0].(float64)
}

func (data TradingPairBook) GetCount() float64 {
	return data[1].(float64)
}

// GetAmount gets total amount available at that price level
// (if AMOUNT > 0 then ask else bid)
func (data TradingPairBook) GetAmount() float64 {
	return data[2].(float64)
}

func (data *FundingCurrencyBooks) FromRaw(raw []byte) error {
	err := json.Unmarshal(raw, data)
	if err != nil {
		return err
	}
	return nil
}

func (data FundingCurrencyBook) GetRate() float64 {
	return data[0].(float64)
}

func (data FundingCurrencyBook) GetPeriod() float64 {
	return data[1].(float64)
}

func (data FundingCurrencyBook) GetCount() float64 {
	return data[2].(float64)
}

// GetAmount gets total amount available at that price level
// (if AMOUNT > 0 then ask else bid)
func (data FundingCurrencyBook) GetAmount() float64 {
	return data[3].(float64)
}
