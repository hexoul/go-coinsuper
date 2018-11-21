package types

// Options for request
type Options struct {
	AccessKey   string `json:"accesskey,omitempty"`
	SecretKey   string `json:"secretkey,omitempty"`
	Size        string `json:"size,omitempty"`
	StartDealNo string `json:"startDealNo,omitempty"`
	Symbol      string `json:"symbol,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	UtcStart    string `json:"utcStart,omitempty"`
	UtcEnd      string `json:"utcEnd,omitempty"`
	WithTrade   string `json:"withTrade,omitempty"`
}

// OptionOrder is used for ordering
// options should be sorted alphabetically with key (a-z)
var OptionOrder = []string{
	"accesskey",
	"secretkey",
	"size",
	"startDealNo",
	"symbol",
	"timestamp",
	"utcStart",
	"utcEnd",
	"withTrade",
}
