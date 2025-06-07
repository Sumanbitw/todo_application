package routes

import (
	"mongoapi/controllers"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/todos", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/api/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/{id}", controllers.CheckTodo).Methods("PUT")
	router.HandleFunc("/api/todos/{id}", controllers.DeleteOneTodo).Methods("DELETE")
	router.HandleFunc("/api/deleteAllTodos", controllers.DeleteManyTodo).Methods("DELETE")

	return router
}
