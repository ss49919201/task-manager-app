package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/s-beats/rest-todo/di"
	"github.com/samber/do"
)

type (
	createUserInput struct {
		Name string `json:"name"`
	}
	updateUserInput struct {
		Name string `json:"name"`
	}

	rowUser struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

func CreateUser(input *HandlerInput) *HandlerOutput {
	var createInput createUserInput
	decoder := json.NewDecoder(input.Body)
	if err := decoder.Decode(&createInput); err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusBadRequest,
			Body:       []byte(string(fmt.Errorf("failed to decode input body: %w", err).Error())),
		}
	}

	db := do.MustInvoke[*sql.DB](di.NewContainer().Injector)
	if _, err := db.QueryContext(input.Context, "INSERT INTO users (id, name) VALUES (?, ?)", uuid.New(), createInput.Name); err != nil {
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

func UpdateUser(input *HandlerInput) *HandlerOutput {
	id := input.URL.Path[len("/users/"):]

	var updateInput updateUserInput
	decoder := json.NewDecoder(input.Body)
	if err := decoder.Decode(&updateInput); err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusBadRequest,
			Body:       []byte(string(fmt.Errorf("failed to decode input body: %w", err).Error())),
		}
	}

	db := do.MustInvoke[*sql.DB](di.NewContainer().Injector)
	tx, err := db.BeginTx(input.Context, nil)
	if err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}

	rows, err := tx.QueryContext(input.Context, "SELECT id, name FROM users FOR UPDATE")
	if err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}
	defer rows.Close()

	user := &rowUser{}
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return &HandlerOutput{
				StatusCode: http.StatusInternalServerError,
				Body:       []byte(string(err.Error())),
			}
		}
	} else {
		return &HandlerOutput{
			StatusCode: http.StatusNotFound,
			Body:       []byte("null"),
		}
	}

	if _, err := tx.QueryContext(input.Context, "UPDATE users SET name = ? WHERE id = ?", updateInput.Name, id); err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}

	if err := tx.Commit(); err != nil {
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
	db := do.MustInvoke[*sql.DB](di.NewContainer().Injector)
	rows, err := db.QueryContext(input.Context, "SELECT id, name FROM users")
	if err != nil {
		return &HandlerOutput{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(string(err.Error())),
		}
	}
	defer rows.Close()

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
		StatusCode: http.StatusOK,
		Body:       b,
	}
}
