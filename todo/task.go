package todo

import "time"

type Task struct {
	ID          int
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func CreateTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}
