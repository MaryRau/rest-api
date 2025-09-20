package http

import (
	"encoding/json"
	"errors"
	"time"
)

type ErrorDTO struct {
	Message string
	Time    time.Time
}

type TaskDTO struct {
	Title       string
	Description string
}

type CompleteTaskDTO struct {
	IsCompleted bool
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (t TaskDTO) CheckForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}
