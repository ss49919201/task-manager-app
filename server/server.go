package server

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/s-beats/rest-todo/di"
	"github.com/s-beats/rest-todo/log"
	"github.com/samber/do"

	"github.com/s-beats/rest-todo/infra/rdb"
	"github.com/s-beats/rest-todo/usecase"
)

func Start() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	return http.ListenAndServe(net.JoinHostPort(host, port), nil)
}

type middleware func(http.HandlerFunc) http.HandlerFunc

type router struct {
	middlewareFunctions []middleware
}

func NewRouter() *router {
	return &router{}
}

// TODO: 重複の対応
func (r *router) SetMiddleware(m middleware) *router {
	if m == nil {
		return r
	}

	r.middlewareFunctions = append(r.middlewareFunctions, m)
	return r
}

func usecaseMiddlewarefunc(f http.HandlerFunc) http.HandlerFunc {
	type contextKey string
	const (
		contextKeyTaskUsecase contextKey = "taskUsecase"
		contextKeyUserUsecase contextKey = "userUsecase"
	)

	return func(w http.ResponseWriter, r *http.Request) {
		db, err := rdb.NewDB()
		if err != nil {
			log.Fatal().Err(err)
		}

		diContainer := di.NewContainer(db)

		taskUsecase := do.MustInvoke[usecase.Task](diContainer.Injector)
		userUsecase := do.MustInvoke[usecase.User](diContainer.Injector)

		m := map[contextKey]any{
			contextKeyTaskUsecase: taskUsecase,
			contextKeyUserUsecase: userUsecase,
		}
		for k, v := range m {
			r = r.WithContext(context.WithValue(r.Context(), k, v))
		}

		f(w, r)
	}
}
