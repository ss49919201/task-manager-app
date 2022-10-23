package persistence

import (
	"context"
	"time"

	"github.com/s-beats/rest-todo/domain"
	"github.com/s-beats/rest-todo/domain/repository"
	"github.com/s-beats/rest-todo/infra/rdb/persistence/internal"
	"golang.org/x/xerrors"
	"xorm.io/xorm"
)

type task struct {
	db *xorm.Engine
}

func NewTask(db *xorm.Engine) repository.Task {
	return &task{
		db: db,
	}
}

type TaskDTO struct {
	ID         string `xorm:"'id'"`
	Title      string
	Text       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     string `xorm:"'user_id'"`
	PriorityID uint   `xorm:"'priority_id'"`
}

func (t *TaskDTO) TableName() string {
	return "tasks"
}

func (t *task) Save(ctx context.Context, task *domain.Task) error {
	taskDTO := &TaskDTO{
		ID:        task.ID().String(),
		Title:     task.Title().String(),
		Text:      task.Text().String(),
		CreatedAt: task.CreatedAt(),
		UpdatedAt: task.UpdatedAt(),
		UserID:    task.CreatedBy().ID().String(),
	}

	var taskPriorityID uint
	has, err := t.db.Table("task_priorities").Cols("id").Where("value = ?", task.Priority()).Get(&taskPriorityID)
	if err != nil {
		return xerrors.Errorf("%v", err)
	}
	if !has {
		return xerrors.New("invalid task priority")
	}

	taskDTO.PriorityID = taskPriorityID

	session := t.db.NewSession().Context(ctx)
	if err := internal.Upsert(session, taskDTO, func(session *xorm.Session) {
		session.
			Cols("title", "text", "updated_at", "priority_id").
			Where("id = ?", taskDTO.ID)
	}); err != nil {
		return err
	}

	return nil
}

func (t *task) GetOne(ctx context.Context, taskID domain.TaskID) (*domain.Task, error) {
	panic("not implement")
}
