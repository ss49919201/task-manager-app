package service

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
	"github.com/s-beats/rest-todo/domain/repository"
	"github.com/s-beats/rest-todo/util"
	"github.com/samber/mo"
)

type Task interface {
	CreateTaskByUser(ctx context.Context, title, text, userID, priority string) mo.Result[*domain.Task]
}

type task struct {
	userRepository repository.User
	taskRepository repository.Task
}

func NewTask(userRepo repository.User, taskRepo repository.Task) Task {
	return &task{
		userRepository: userRepo,
		taskRepository: taskRepo,
	}
}

func (t *task) CreateTaskByUser(ctx context.Context, title, text, userID, priority string) mo.Result[*domain.Task] {
	result := t.userRepository.GetOne(ctx, domain.NewUserID(userID))
	if result.Error() != nil {
		return domain.ToErrTask(result.Error())
	}
	user := result.MustGet()

	now := util.GetTimeNow()
	return domain.NewTask(
		domain.NewTaskID(util.NewUUID()),
		domain.NewTaskTitle(title),
		domain.NewTaskText(text),
		now,
		now,
		user,
		domain.NewPriority(priority),
	)
}
