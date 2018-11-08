package types

// Options for request
type Options struct {
	Symbol      string `json:"symbol,omitempty"`
	UtcStart    string `json:"utcStart,omitempty"`
	UtcEnd      string `json:"utcEnd,omitempty"`
	StartDealNo string `json:"startDealNo,omitempty"`
	Size        string `json:"size,omitempty"`
}
