package model

type Response struct {
	Success    bool        `json:"success"`
	StatusCode *int        `json:"status_code,omitempty"`
	Message    *string     `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
