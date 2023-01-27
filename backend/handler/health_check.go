package handler

import "net/http"

func HealthCheck(_ *HandlerInput) *HandlerOutput {
	return &HandlerOutput{
		StatusCode: http.StatusOK,
		Body:       []byte("OK"),
	}
}
