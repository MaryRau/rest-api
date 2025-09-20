package database

import (
	"context"
	"restapi/todo"
	"time"
)

// добавление задания в базу данных
func AddTask(task todo.Task) error {
	_, err := DB.Exec(context.Background(), "INSERT INTO Todos (title, description, isCompleted, createdAt) VALUES ($1, $2, $3, $4)", task.Title, task.Description, task.IsCompleted, task.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// отметка о выполнении задания
func CompleteTask(id int) (todo.Task, error) {
	now := time.Now()
	_, err := DB.Exec(context.Background(), "UPDATE Todos SET isCompleted = TRUE, completedAt = $1 WHERE id = $2", &now, id)
	if err != nil {
		return todo.CreateTask("", ""), err
	}

	row, err := DB.Query(context.Background(), "SELECT * FROM Todos WHERE id = $1", id)
	if err != nil {
		return todo.CreateTask("", ""), err
	}

	defer row.Close()

	var task todo.Task
	for row.Next() {
		if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.CompletedAt); err != nil {
			return todo.CreateTask("", ""), err
		}
	}

	return task, nil
}

// удаление задания из базы данных
func DeleteTask(id int) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM Todos WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

// получение всех заданий
func GetAllTasks() ([]todo.Task, error) {
	rows, err := DB.Query(context.Background(), "SELECT * FROM Todos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []todo.Task

	for rows.Next() {
		var task todo.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.CompletedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// получение выполненных/невыполненных заданий
func GetTasksByCompliting(isCompleted bool) ([]todo.Task, error) {
	rows, err := DB.Query(context.Background(), "SELECT * FROM Todos WHERE isCompleted = $1", isCompleted)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []todo.Task

	for rows.Next() {
		var task todo.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted, &task.CreatedAt, &task.CompletedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
