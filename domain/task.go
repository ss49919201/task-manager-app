package domain

import (
	"errors"
	"time"

	"github.com/samber/mo"
)

type Task struct {
	id        TaskID
	title     TaskTitle
	text      TaskText
	createdAt time.Time
	updatedAt time.Time
	createdBy *User
	priority  Priority
}

func NewTask(id TaskID, title TaskTitle, text TaskText, createdAt, updatedAt time.Time, createdBy *User, priority Priority) mo.Result[*Task] {
	if createdAt.After(updatedAt) {
		return ToErrTask(errors.New("createdAt must be LTE updatedAt"))
	}

	return ToOKTask(
		&Task{
			id:        id,
			title:     title,
			text:      text,
			createdAt: createdAt,
			updatedAt: updatedAt,
			createdBy: createdBy,
			priority:  priority,
		},
	)
}

func (t *Task) ID() *TaskID {
	return t.id
}

func (t *Task) Title() *TaskTitle {
	return t.title
}

func (t *Task) Text() *TaskText {
	return t.text
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) CreatedBy() *User {
	return t.createdBy
}

func (t *Task) Priority() Priority {
	return t.priority
}

func ToOKTask(v *Task) mo.Result[*Task] {
	return mo.Ok(v)
}

func ToErrTask(err error) mo.Result[*Task] {
	return mo.Err[*Task](err)
}
