package handler

import (
	"context"
	"io"
	"net/http"
)

type HandlerInput struct {
	context.Context
	Body io.Reader
}

type HandlerOutput struct {
	StatusCode int
	Body       []byte
}

func Wrap(handler func(input *HandlerInput) *HandlerOutput) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		output := handler(&HandlerInput{
			Context: r.Context(),
			Body:    r.Body,
		})
		w.WriteHeader(output.StatusCode)
		w.Write(output.Body)
	}
}
