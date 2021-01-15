package models

type Response struct {
	Success bool `json:"success"`
	Error interface{} `json:"error,omitempty"`
	Data interface{} `json:"data,omitempty"`
}