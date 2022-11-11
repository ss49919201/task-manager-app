package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/s-beats/rest-todo/infra/rdb"
)

var db *sql.DB

func lazyInitDB() *sql.DB {
	if db == nil {
		var err error
		db, err = rdb.NewDB()
		if err != nil {
			panic(err)
		}
	}
	return db
}

func CreateUser(input *HandlerInput) *HandlerOutput {
	type userInput struct {
		Name string `json:"name"`
	}

	var u userInput
	decoder := json.NewDecoder(input.Body)
	if err := decoder.Decode(&u); err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusBadRequest,
			Body:       []byte(string(fmt.Errorf("failed to decode input body: %w", err).Error())),
		}
	}

	db := lazyInitDB()
	if _, err := db.QueryContext(input.Context, "INSERT INTO users (id, name) VALUES (?, ?)", uuid.New(), u.Name); err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}

	return &HandlerOutput{
		StatusCode: http.StatusCreated,
		Body:       []byte("null"),
	}
}

func GetUserList(input *HandlerInput) *HandlerOutput {
	db := lazyInitDB()
	rows, err := db.QueryContext(input.Context, "SELECT id, name FROM users")
	if err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}
	defer rows.Close()

	type rowUser struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	users := []*rowUser{}
	for rows.Next() {
		u := &rowUser{}
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return &HandlerOutput{
				StatusCode: http.StatusInternalServerError,
				Body:       []byte(string(err.Error())),
			}
		}
		users = append(users, u)
	}

	b, err := json.Marshal(users)
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
