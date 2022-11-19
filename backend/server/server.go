package server

import (
	"net"
	"net/http"
	"os"

	"github.com/s-beats/rest-todo/handler"
	"github.com/s-beats/rest-todo/log"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Router interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	GET(pattern string, fn http.HandlerFunc)
	POST(pattern string, fn http.HandlerFunc)
	PushBackMiddleware(m Middleware) Router
}

func defineRoutes(router Router) {
	router.GET("/users", handler.Wrap(handler.GetUserList))
	router.POST("/users", handler.Wrap(handler.CreateUser))
}

func Start(router Router) error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	defineRoutes(router)

	log.Info().Msgf("Starting server on %s:%s", host, port)
	return http.ListenAndServe(net.JoinHostPort(host, port), router)
}
