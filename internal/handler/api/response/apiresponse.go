// Package response provides HTTP response utilities and structures
package response

import (
	"encoding/json"
	"net/http"
)

// APIResponse represents a generic API response with data
type APIResponse[T any] struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
	Data       *T     `json:"data,omitempty"`
	Error      string `json:"error,omitempty"`
}

// APIResponseNoData represents an API response without data
type APIResponseNoData struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty"`
}

// Success creates a successful API response with data
func Success[T any](data T, msg string) APIResponse[T] {
	return APIResponse[T]{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    msg,
		Data:       &data,
	}
}

// Fail creates a failed API response with data
func Fail[T any](data T, statusCode int, err, msg string) APIResponse[T] {
	return APIResponse[T]{
		StatusCode: statusCode,
		Success:    false,
		Message:    msg,
		Error:      err,
		Data:       &data,
	}
}

// SuccessNoData creates a successful API response without data
func SuccessNoData(msg string) APIResponseNoData {
	return APIResponseNoData{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    msg,
	}
}

// FailNoData creates a failed API response without data
func FailNoData(status int, err, msg string) APIResponseNoData {
	return APIResponseNoData{
		StatusCode: status,
		Success:    false,
		Error:      err,
		Message:    msg,
	}
}

// ----------------------------
// gRPC Error Handler
// ----------------------------

// ----------------------------
// Write JSON Response
// ----------------------------

// WriteSuccessNoData writes a successful response without data to the HTTP response writer
func WriteSuccessNoData(w http.ResponseWriter, msg string) {
	res := SuccessNoData(msg)
	writeJSON(w, res.StatusCode, res)
}

// WriteSuccess writes a successful response with data to the HTTP response writer
func WriteSuccess[T any](w http.ResponseWriter, data T, msg string) {
	res := Success(data, msg)
	writeJSON(w, res.StatusCode, res)
}

// WriteFailNoData writes a failed response without data to the HTTP response writer
func WriteFailNoData(w http.ResponseWriter, statusCode int, err string, msg string) {
	res := FailNoData(statusCode, err, msg)
	writeJSON(w, res.StatusCode, res)
}

// WriteFailWithData writes a failed response with data to the HTTP response writer
func WriteFailWithData[T any](w http.ResponseWriter, data T, statusCode int, err string, msg string) {
	res := Fail(data, statusCode, err, msg)
	writeJSON(w, res.StatusCode, res)
}

func writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) {
	out, err := json.Marshal(data)
	if err != nil {
		return
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(out)
}
