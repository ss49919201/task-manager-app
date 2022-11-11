package usecase

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
	"github.com/s-beats/rest-todo/domain/repository"
	"github.com/s-beats/rest-todo/domain/service"
	"github.com/s-beats/rest-todo/util"
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
	wappedSave := util.ConvertMapperWithCtx(ctx, t.taskRepository.Save)

	result := t.taskService.CreateTaskByUser(ctx, title, text, userID, priority).FlatMap(wappedSave)
	if result.Error() != nil {
		return nil, result.Error()
	}

	return result.MustGet(), nil
}
