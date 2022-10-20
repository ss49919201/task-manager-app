package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
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

func createUser(u usecase.User) error {
	panic("not implemented")
}

func createTask(u usecase.Task) error {
	panic("not implemented")
}

func main() {
	db, err := rdb.NewDB()
	if err != nil {
		log.Fatal().Err(err)
	}

	taskRepo := repository.NewTask(db)
	userRepo := repository.NewUser(db)
	_ = usecase.NewTask(taskRepo, userRepo)
	userUsecase := usecase.NewUser(userRepo)

	if err := createUser(userUsecase); err != nil {
		log.Error().Err(err).Send()
	}
}
