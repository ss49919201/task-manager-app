package domain

import (
	"time"

	"github.com/s-beats/rest-todo/util"
	"github.com/samber/mo"
)

type Task struct {
	id        *TaskID
	title     *TaskTitle
	text      *TaskText
	createdAt time.Time
	updatedAt time.Time
	createdBy *User
	priority  Priority
}

func NewTask(id *TaskID, title *TaskTitle, text *TaskText, createdAt, updatedAt time.Time, createdBy *User, priority Priority) mo.Result[*Task] {
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

type TaskID struct {
	id string
}

func (t *TaskID) String() string {
	return t.id
}

func NewTaskID(id string) *TaskID {
	return &TaskID{
		id: util.UUIDMustParse(id),
	}
}

type TaskTitle struct {
	title string
}

func (t *TaskTitle) String() string {
	return t.title
}

func NewTaskTitle(title string) *TaskTitle {
	return &TaskTitle{
		title: title,
	}
}

type TaskText struct {
	text string
}

func (t *TaskText) String() string {
	return t.text
}

func NewTaskText(text string) *TaskText {
	return &TaskText{
		text: text,
	}
}

func ToOKTask(v *Task) mo.Result[*Task] {
	return mo.Ok(v)
}

func ToErrTask(err error) mo.Result[*Task] {
	return mo.Err[*Task](err)
}
