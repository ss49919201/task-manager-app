package server

import (
	"context"
	"net/http"

	"github.com/s-beats/rest-todo/di"
	"github.com/s-beats/rest-todo/infra/rdb"
	"github.com/s-beats/rest-todo/log"
	"github.com/s-beats/rest-todo/usecase"
	"github.com/samber/do"
)

func usecaseMiddlewarefunc(f http.HandlerFunc) http.HandlerFunc {
	type contextKey string
	const (
		contextKeyTaskUsecase contextKey = "taskUsecase"
		contextKeyUserUsecase contextKey = "userUsecase"
	)

	return func(w http.ResponseWriter, r *http.Request) {
		db, err := rdb.NewDB()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to database")
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
