package logger

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Response struct for standardized API responses
type Response struct {
	Code string      `json:"response_code"`
	Desc string      `json:"response_desc"`
	Data interface{} `json:"response_data"`
}

// ResponseKey for storing response in context
type ResponseKey string
type ContextKey string

const (
	LogFieldsKey   ContextKey  = "logFields"
	ResponseCtxKey ResponseKey = "responseData"
)

// WithLogFields menambahkan log fields ke context
func WithLogFields(ctx context.Context, fields logrus.Fields) context.Context {
	return context.WithValue(ctx, LogFieldsKey, fields)
}

// GetLogFields mengambil log fields dari context
func GetLogFields(ctx context.Context) logrus.Fields {
	if fields, ok := ctx.Value(LogFieldsKey).(logrus.Fields); ok {
		return fields
	}
	return logrus.Fields{}
}

// InfoWithContext mencatat log info dengan context
func InfoWithContext(ctx context.Context, message string) {
	fields := GetLogFields(ctx)
	Info(message, fields)
}

// ErrorWithContext mencatat log error dengan context
func ErrorWithContext(ctx context.Context, message string, err error) {
	fields := GetLogFields(ctx)
	Error(message, err, fields)
}

// WithResponse adds a Response to the context
func WithResponse(ctx context.Context, response Response) context.Context {
	return context.WithValue(ctx, ResponseCtxKey, response)
}

// GetResponse retrieves the Response from the context
func GetResponse(ctx context.Context) (Response, bool) {
	response, ok := ctx.Value(ResponseCtxKey).(Response)
	return response, ok
}

// LogAndSendResponse logs and sends the response
func LogAndSendResponse(ctx context.Context, w http.ResponseWriter) {
	response, ok := GetResponse(ctx)
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logrus.Error("No response found in context")
		return
	}

	// Log the response
	fields := GetLogFields(ctx)
	log.WithFields(fields).Info("Sending response", logrus.Fields{
		"response_code": response.Code,
		"response_desc": response.Desc,
	})

	// Send response to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
