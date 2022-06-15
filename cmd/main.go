package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
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
	log.Fatal().Err(server.Start()).Send()
}
