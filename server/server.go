package server

import (
	"net"
	"net/http"
	"os"

	"github.com/s-beats/rest-todo/log"
)

func Start() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	log.Info().Msgf("Starting server on %s:%s", host, port)
	return http.ListenAndServe(
		net.JoinHostPort(host, port),
		NewRouter(),
	)
}
