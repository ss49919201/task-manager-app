package handler

import "net/http"

func CreateUser(input *HandlerInput) *HandlerOutput {
	// TODO
	return &HandlerOutput{
		StatusCode: http.StatusCreated,
		Body:       []byte("{\"id\": 1, \"name\": \"John\"}"),
	}
}

func GetUserList(input *HandlerInput) *HandlerOutput {
	// TODO
	return &HandlerOutput{
		StatusCode: http.StatusOK,
		Body:       []byte("[{\"id\": 1, \"name\": \"John\"}]"),
	}
}
