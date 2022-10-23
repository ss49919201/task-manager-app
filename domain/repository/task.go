package repository

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
)

type Task interface {
	Save(ctx context.Context, task *domain.Task) error
	GetOne(ctx context.Context, taskID domain.TaskID) (*domain.Task, error)
}
