package utils

import "time"

// Response is the standard API response structure
type Response struct {
	Status    string      `json:"status"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// SuccessResponse creates a success response
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:    "success",
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// ErrorResponse creates an error response
func ErrorResponse(status string, error string) Response {
	return Response{
		Status:    status,
		Error:     error,
		Timestamp: time.Now(),
	}
}
