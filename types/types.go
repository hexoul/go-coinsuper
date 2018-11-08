// Package types of both request and response for API
package types

// ResponseData structure
type ResponseData struct {
	Timestamp string      `json:"timestamp"`
	Result    interface{} `json:"result"`
}

// Response structure
type Response struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data *ResponseData `json:"data"`
}

// RequestCommon structure
type RequestCommon struct {
	AccessKey string `json:"accesskey"`
	Timestamp string `json:"timestamp,omitempty"`
	Sign      string `json:"sign"`
}

// RequestCommon structure
type Request struct {
	Common *RequestCommon `json:"common"`
	Data   *Options       `json:"data"`
}
