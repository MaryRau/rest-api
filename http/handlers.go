package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/database"
	"restapi/todo"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// запрос на создание задания
/*
pattern: /task
method:  POST
info:	 JSON in request body,
		 JSON in response body
*/
func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.CheckForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	newTask := todo.CreateTask(taskDTO.Title, taskDTO.Description)
	if err := database.AddTask(newTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(newTask, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response:", err)
	}
}

// запрос на получение всех заданий
/*
pattern: /task
method:  GET
info:	 JSON in response body
*/
func HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := database.GetAllTasks()
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response:", err)
	}
}

// запрос на получение выполненных/невыполненных заданий
/*
pattern: /task&completed=true/false
method:  GET
info:	 JSON in response body
*/
func HandleGetTasksByCompliting(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if query.Get("completed") == "" {
		http.Error(w, "Не выделяется значение выполнения", http.StatusInternalServerError)
		return
	}

	compliting, err := strconv.ParseBool(query.Get("completed"))
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	tasks, err := database.GetTasksByCompliting(compliting)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response:", err)
	}
}

// запрос на выполнение задания по его id
/*
pattern: /task/id
method:  PATCH
info:	 JSON in request body,
		 JSON in response body
*/
func HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeTaskDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&completeTaskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	task, err := database.CompleteTask(id)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write http response:", err)
	}
}

// запрос на удаление задания по его id
/*
pattern: /task/id
method:  DELETE
info:	 no info
*/
func HandleDeteleTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	if err := database.DeleteTask(id); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
