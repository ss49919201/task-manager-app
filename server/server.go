package server

import (
	"net"
	"net/http"
	"os"

	"github.com/s-beats/rest-todo/handler"
	"github.com/s-beats/rest-todo/log"
)

func defineRoutes(router *router) {
	router.Get("/users", handler.Wrap(handler.GetUserList))
	router.Post("/users", handler.Wrap(handler.GetUserList))
}

func Start() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	router := NewRouter()
	defineRoutes(router)

	log.Info().Msgf("Starting server on %s:%s", host, port)
	return http.ListenAndServe(net.JoinHostPort(host, port), router)
}
