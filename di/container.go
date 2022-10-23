package di

import (
	"sync"

	"github.com/s-beats/rest-todo/domain/repository"
	"github.com/s-beats/rest-todo/domain/service"
	"github.com/s-beats/rest-todo/infra/rdb/persistence"
	"github.com/s-beats/rest-todo/usecase"
	"github.com/samber/do"
	"xorm.io/xorm"
)

type Container struct {
	*do.Injector
}

var (
	container *Container
	syncOnce  sync.Once
)

func NewContainer(db *xorm.Engine) *Container {
	syncOnce.Do(func() {
		container = &Container{
			do.New(),
		}

		// repository
		do.ProvideValue(container.Injector, persistence.NewUser(db))
		do.ProvideValue(container.Injector, persistence.NewTask(db))

		// service
		do.ProvideValue(container.Injector, service.NewTask(
			do.MustInvoke[repository.User](container.Injector),
			do.MustInvoke[repository.Task](container.Injector),
		))

		// usecase
		do.ProvideValue(container.Injector, usecase.NewTask(
			do.MustInvoke[service.Task](container.Injector),
			do.MustInvoke[repository.Task](container.Injector),
		))
		do.ProvideValue(container.Injector, usecase.NewUser(
			do.MustInvoke[repository.User](container.Injector),
		))
	})

	return container
}
