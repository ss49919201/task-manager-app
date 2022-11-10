package main

import (
	"github.com/joho/godotenv"
	"github.com/s-beats/rest-todo/log"
	"github.com/s-beats/rest-todo/router"
	"github.com/s-beats/rest-todo/server"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg(("Error loading env file"))
	}
}

func main() {
	log.Fatal().
		Err(server.Start(router.NewRouter())).
		Send()
}
