package models

type EchoResponse struct {
	Message    string `json:"msg"`
	Timestramp string `json:"timestamp"`
}
