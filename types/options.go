package types

// Options for request
type Options struct {
	Size        string `json:"size,omitempty"`
	StartDealNo string `json:"startDealNo,omitempty"`
	Symbol      string `json:"symbol,omitempty"`
	UtcStart    string `json:"utcStart,omitempty"`
	UtcEnd      string `json:"utcEnd,omitempty"`
	WithTrade   string `json:"withTrade,omitempty"`
}
