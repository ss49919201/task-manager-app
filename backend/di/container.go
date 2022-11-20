package di

import (
	"github.com/s-beats/rest-todo/infra"
	"github.com/s-beats/rest-todo/log"
	"github.com/samber/do"
)

func init() {
	container = &Container{
		do.New(),
	}

	db, err := infra.NewDB()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	// db
	do.ProvideValue(container.Injector, db)
}

type Container struct {
	*do.Injector
}

var (
	container *Container
)

func NewContainer() *Container {
	return &Container{
		container.Clone(),
	}
}
