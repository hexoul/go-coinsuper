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
