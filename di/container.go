package di

import (
	"sync"

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

		// db
		do.ProvideValue(container.Injector, db)
	})

	return container
}
