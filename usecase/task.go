package usecase

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
	"github.com/s-beats/rest-todo/domain/repository"
	"github.com/s-beats/rest-todo/util"
	"github.com/samber/mo"
)

type Task interface {
	Create(ctx context.Context, title, text, userID, priority string) (*domain.Task, error)
}

type task struct {
	taskRepository repository.Task
	userRepository repository.User
}

func NewTask(taskRepo repository.Task, userRepo repository.User) Task {
	return &task{
		taskRepository: taskRepo,
		userRepository: userRepo,
	}
}

func (t *task) Create(ctx context.Context, title, text, userID, priority string) (*domain.Task, error) {
	result1 := t.userRepository.GetOne(ctx, *domain.NewUserID(userID))
	if result1.Error() != nil {
		return nil, result1.Error()
	}
	user := result1.MustGet()

	now := util.GetTimeNow()
	result2 := domain.NewTask(
		domain.NewTaskID(util.NewUUID()),
		domain.NewTaskTitle(title),
		domain.NewTaskText(text),
		now,
		now,
		user,
		domain.NewPriority(priority),
	).FlatMap(func(val *domain.Task) mo.Result[*domain.Task] {
		return t.taskRepository.Save(ctx, val)
	})
	if result2.Error() != nil {
		return nil, result2.Error()
	}

	return result2.MustGet(), nil
}
