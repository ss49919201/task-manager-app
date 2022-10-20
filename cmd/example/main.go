package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/s-beats/rest-todo/domain"
	"github.com/s-beats/rest-todo/domain/repository"
	"github.com/s-beats/rest-todo/infra/rdb"
	"github.com/s-beats/rest-todo/usecase"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg(("Error loading env file"))
	}
}

func exec[T any](fn func() (T, error)) {
	// TODO: 実装
}

func createUser(u usecase.User) (*domain.User, error) {
	return u.Create(context.Background(), "username")
}

func createTask(u usecase.Task, userID string) (*domain.Task, error) {
	return u.Create(context.Background(), "title", "description", userID, "HIGH")
}

func main() {
	db, err := rdb.NewDB()
	if err != nil {
		log.Fatal().Err(err)
	}

	taskRepo := repository.NewTask(db)
	userRepo := repository.NewUser(db)
	taskUsecase := usecase.NewTask(taskRepo, userRepo)
	userUsecase := usecase.NewUser(userRepo)

	user, err := createUser(userUsecase)
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	// 本当`Interfacee()`にしたいが、Privateなフィールドは出力できないので`Str()`
	log.Info().Interface("action", "createUser").Interface("userID", user.ID().String()).Send()

	task, err := createTask(taskUsecase, user.ID().String())
	if err != nil {
		log.Error().Err(err).Send()
		return
	}

	log.Info().Interface("action", "createTask").Interface("taskID", task.ID().String()).Send()
}
