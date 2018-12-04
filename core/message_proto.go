package core

// Message is a type or struct for Messaging API
type Message struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// Result is a type or struct for messaging API triggered status
type Result struct {
	Trigger string `json:"trigger"`
	Message string `json:"message"`
	Result  string `json:"result"` // success or failed
	Data    string `json:"data"`
	Error   string `json:"error"`
	Detail  string `json:"error_detail"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Tip     string `json:tip`
}
