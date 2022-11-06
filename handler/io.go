package handler

import "net/http"

type HandlerInput struct{}

type HandlerOutput struct{}

func Wrap(handler func(input *HandlerInput) *HandlerOutput) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
	}
}
