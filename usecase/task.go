package usecase

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
	"github.com/s-beats/rest-todo/domain/repository"
	"github.com/s-beats/rest-todo/domain/service"
	"github.com/samber/mo"
)

type Task interface {
	Create(ctx context.Context, title, text, userID, priority string) (*domain.Task, error)
}

type task struct {
	taskService    service.Task
	taskRepository repository.Task
}

func NewTask(taskService service.Task, taskRepository repository.Task) Task {
	return &task{
		taskService:    taskService,
		taskRepository: taskRepository,
	}
}

func (t *task) Create(ctx context.Context, title, text, userID, priority string) (*domain.Task, error) {
	result := t.taskService.CreateTaskByUser(ctx, title, text, userID, priority).FlatMap(a(ctx, t.taskRepository.Save))
	if result.Error() != nil {
		return nil, result.Error()
	}

	return result.MustGet(), nil
}

func a[T any](ctx context.Context, fn func(ctx context.Context, val T) mo.Result[T]) func(v T) mo.Result[T] {
	return func(v T) mo.Result[T] {
		return fn(ctx, v)
	}
}
