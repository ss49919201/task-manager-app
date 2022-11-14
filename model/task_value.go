package model

import (
	"github.com/s-beats/rest-todo/util"
)

type TaskID struct {
	id string
}

func (t TaskID) String() string {
	return t.id
}

func NewTaskID(id string) TaskID {
	return TaskID{
		id: util.UUIDMustParse(id),
	}
}

type TaskTitle struct {
	title string
}

func (t TaskTitle) String() string {
	return t.title
}

func NewTaskTitle(title string) TaskTitle {
	return TaskTitle{
		title: title,
	}
}

type TaskText struct {
	text string
}

func (t TaskText) String() string {
	return t.text
}

func NewTaskText(text string) TaskText {
	return TaskText{
		text: text,
	}
}
