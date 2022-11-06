package handler

import (
	"encoding/json"
	"net/http"

	"github.com/s-beats/rest-todo/infra/rdb"
)

func CreateUser(input *HandlerInput) *HandlerOutput {
	// TODO
	db, err := rdb.NewDB()
	if err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}

	type userInput struct {
		Name string `json:"name"`
	}

	var u userInput
	if err := json.Unmarshal(input.Body, &u); err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusBadRequest,
			Body:       []byte(string(err.Error())),
		}
	}

	rows, err := db.QueryContext(input.Context, "INSERT INTO users (name) VALUES (?)", u.Name)
	if err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}
	defer rows.Close()

	type row struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var r row
	if rows.Next() {
		if err := rows.Scan(&r.ID, &r.Name); err != nil {
			return &HandlerOutput{
				StatusCode: http.StatusInternalServerError,
				Body:       []byte(string(err.Error())),
			}
		}
	}

	b, err := json.Marshal(r)
	if err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}

	return &HandlerOutput{
		StatusCode: http.StatusCreated,
		Body:       b,
	}
}

func GetUserList(input *HandlerInput) *HandlerOutput {
	// TODO
	return &HandlerOutput{
		StatusCode: http.StatusOK,
		Body:       []byte("[{\"id\": 1, \"name\": \"John\"}]"),
	}
}
