package dto

import "time"

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type HealthData struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
