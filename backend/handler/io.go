package handler

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type HandlerInput struct {
	context.Context
	URL  *url.URL
	Body io.Reader
}

type HandlerOutput struct {
	StatusCode int
	Body       []byte
}

func Wrap(handler func(input *HandlerInput) *HandlerOutput) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
		input := &HandlerInput{
			Context: r.Context(),
			URL:     r.URL,
			Body:    r.Body,
		}
		output := handler(input)
		w.WriteHeader(output.StatusCode)
		w.Write(output.Body)
	}
}
