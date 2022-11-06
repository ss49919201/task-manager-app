package handler

import (
	"context"
	"net/http"
)

type HandlerInput struct {
	context.Context
	Body []byte
}

type HandlerOutput struct {
	StatusCode int
	Body       []byte
}

func Wrap(handler func(input *HandlerInput) *HandlerOutput) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		output := handler(&HandlerInput{})
		w.WriteHeader(output.StatusCode)
		w.Write(output.Body)
	}
}
