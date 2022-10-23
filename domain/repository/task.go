package repository

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
	"github.com/samber/mo"
)

type Task interface {
	Save(ctx context.Context, task *domain.Task) mo.Result[*domain.Task]
	GetOne(ctx context.Context, taskID domain.TaskID) mo.Result[*domain.Task]
}
