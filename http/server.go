package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() error {
	router := mux.NewRouter()

	router.Path("/task").Methods("POST").HandlerFunc(HandleCreateTask)
	router.Path("/task").Methods("GET").Queries("completed", "{completed}").HandlerFunc(HandleGetTasksByCompliting)
	router.Path("/task").Methods("GET").HandlerFunc(HandleGetAllTasks)
	router.Path("/task/{id}").Methods("PATCH").HandlerFunc(HandleCompleteTask)
	router.Path("/task/{id}").Methods("DELETE").HandlerFunc(HandleDeteleTask)

	err := http.ListenAndServe(":9091", router)
	if err != nil {
		return err
	}

	return nil
}
