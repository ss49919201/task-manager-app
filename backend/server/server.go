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

func defineMiddlewares(router Router) {
	router.PushBackMiddleware(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			next(w, r)
		}
	})
}

func defineRoutes(router Router) {
	router.GET("/users", handler.Wrap(handler.GetUserList))
	router.POST("/users", handler.Wrap(handler.CreateUser))
}

func Start(router Router) error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	defineMiddlewares(router)
	defineRoutes(router)

	log.Info().Msgf("Starting server on %s:%s", host, port)
	return http.ListenAndServe(net.JoinHostPort(host, port), router)
}
