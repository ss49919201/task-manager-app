package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	type args struct {
		id        TaskID
		title     TaskTitle
		text      TaskText
		createdAt time.Time
		updatedAt time.Time
		createdBy *User
		priority  Priority
	}
	tests := []struct {
		name string
		args args
		want mo.Result[*Task]
	}{
		{
			name: "正常系",
			args: args{
				id:        TaskID{"1"},
				title:     TaskTitle{"title"},
				text:      TaskText{"text"},
				createdAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				updatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				createdBy: &User{
					id:   UserID{"1"},
					name: "name",
				},
				priority: Priority(1),
			},
			want: mo.Ok(
				&Task{
					id:        TaskID{"1"},
					title:     TaskTitle{"title"},
					text:      TaskText{"text"},
					createdAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					updatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					createdBy: &User{
						id:   UserID{"1"},
						name: "name",
					},
					priority: Priority(1),
				}),
		},
		{
			name: "作成日時が更新日時より後の場合はエラー",
			args: args{
				id:        TaskID{"1"},
				title:     TaskTitle{"title"},
				text:      TaskText{"text"},
				createdAt: time.Date(2020, 1, 1, 0, 0, 0, 1, time.UTC),
				updatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				createdBy: &User{
					id:   UserID{"1"},
					name: "name",
				},
				priority: Priority(1),
			},
			want: mo.Err[*Task](ErrNewTaskInvalidDate),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTask(tt.args.id, tt.args.title, tt.args.text, tt.args.createdAt, tt.args.updatedAt, tt.args.createdBy, tt.args.priority)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTask_ID(t *testing.T) {
	type fields struct {
		id        TaskID
		title     TaskTitle
		text      TaskText
		createdAt time.Time
		updatedAt time.Time
		createdBy *User
		priority  Priority
	}
	tests := []struct {
		name   string
		fields fields
		want   TaskID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				id:        tt.fields.id,
				title:     tt.fields.title,
				text:      tt.fields.text,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
				createdBy: tt.fields.createdBy,
				priority:  tt.fields.priority,
			}
			if got := tr.ID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Title(t *testing.T) {
	type fields struct {
		id        TaskID
		title     TaskTitle
		text      TaskText
		createdAt time.Time
		updatedAt time.Time
		createdBy *User
		priority  Priority
	}
	tests := []struct {
		name   string
		fields fields
		want   *TaskTitle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				id:        tt.fields.id,
				title:     tt.fields.title,
				text:      tt.fields.text,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
				createdBy: tt.fields.createdBy,
				priority:  tt.fields.priority,
			}
			if got := tr.Title(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Text(t *testing.T) {
	type fields struct {
		id        TaskID
		title     TaskTitle
		text      TaskText
		createdAt time.Time
		updatedAt time.Time
		createdBy *User
		priority  Priority
	}
	tests := []struct {
		name   string
		fields fields
		want   *TaskText
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				id:        tt.fields.id,
				title:     tt.fields.title,
				text:      tt.fields.text,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
				createdBy: tt.fields.createdBy,
				priority:  tt.fields.priority,
			}
			if got := tr.Text(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_CreatedAt(t *testing.T) {
	type fields struct {
		id        TaskID
		title     TaskTitle
		text      TaskText
		createdAt time.Time
		updatedAt time.Time
		createdBy *User
		priority  Priority
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				id:        tt.fields.id,
				title:     tt.fields.title,
				text:      tt.fields.text,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
				createdBy: tt.fields.createdBy,
				priority:  tt.fields.priority,
			}
			if got := tr.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_UpdatedAt(t *testing.T) {
	type fields struct {
		id        TaskID
		title     TaskTitle
		text      TaskText
		createdAt time.Time
		updatedAt time.Time
		createdBy *User
		priority  Priority
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				id:        tt.fields.id,
				title:     tt.fields.title,
				text:      tt.fields.text,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
				createdBy: tt.fields.createdBy,
				priority:  tt.fields.priority,
			}
			if got := tr.UpdatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.UpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_CreatedBy(t *testing.T) {
	type fields struct {
		id        TaskID
		title     TaskTitle
		text      TaskText
		createdAt time.Time
		updatedAt time.Time
		createdBy *User
		priority  Priority
	}
	tests := []struct {
		name   string
		fields fields
		want   *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				id:        tt.fields.id,
				title:     tt.fields.title,
				text:      tt.fields.text,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
				createdBy: tt.fields.createdBy,
				priority:  tt.fields.priority,
			}
			if got := tr.CreatedBy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.CreatedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Priority(t *testing.T) {
	type fields struct {
		id        TaskID
		title     TaskTitle
		text      TaskText
		createdAt time.Time
		updatedAt time.Time
		createdBy *User
		priority  Priority
	}
	tests := []struct {
		name   string
		fields fields
		want   Priority
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				id:        tt.fields.id,
				title:     tt.fields.title,
				text:      tt.fields.text,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
				createdBy: tt.fields.createdBy,
				priority:  tt.fields.priority,
			}
			if got := tr.Priority(); got != tt.want {
				t.Errorf("Task.Priority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskID_String(t *testing.T) {
	type fields struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TaskID{
				id: tt.fields.id,
			}
			if got := tr.String(); got != tt.want {
				t.Errorf("TaskID.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTaskID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want TaskID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskTitle_String(t *testing.T) {
	type fields struct {
		title string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TaskTitle{
				title: tt.fields.title,
			}
			if got := tr.String(); got != tt.want {
				t.Errorf("TaskTitle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTaskTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want *TaskTitle
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskTitle(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskText_String(t *testing.T) {
	type fields struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TaskText{
				text: tt.fields.text,
			}
			if got := tr.String(); got != tt.want {
				t.Errorf("TaskText.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTaskText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want TaskText
	}{
		{
			"success",
			args{"test"},
			TaskText{"test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTaskText(tt.args.text)
			assert.Equal(t, got, tt.want)
		})
	}
}
