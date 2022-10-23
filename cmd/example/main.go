package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/s-beats/rest-todo/di"
	"github.com/s-beats/rest-todo/domain"
	"github.com/s-beats/rest-todo/infra/rdb"
	"github.com/s-beats/rest-todo/log"
	"github.com/s-beats/rest-todo/usecase"
	"github.com/samber/do"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg(("Error loading env file"))
	}
}

func run() error {
	db, err := rdb.NewDB()
	if err != nil {
		return err
	}

	diContainer := di.NewContainer(db)

	taskUsecase := do.MustInvoke[usecase.Task](diContainer.Injector)
	userUsecase := do.MustInvoke[usecase.User](diContainer.Injector)

	user, err := createUser(userUsecase)
	if err != nil {
		return err
	}

	// 本当`Interfacee()`にしたいが、Privateなフィールドは出力できないので`Str()`
	log.Info().Interface("action", "createUser").Interface("userID", user.ID().String()).Send()

	task, err := createTask(taskUsecase, user.ID().String())
	if err != nil {
		return err
	}

	log.Info().Interface("action", "createTask").Interface("taskID", task.ID().String()).Send()
	return nil
}

func createUser(u usecase.User) (*domain.User, error) {
	return u.Create(context.Background(), "username")
}

func createTask(u usecase.Task, userID string) (*domain.Task, error) {
	return u.Create(context.Background(), "title", "description", userID, "HIGH")
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err)
	}
}
