package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pri/todos/Controllers"

	"github.com/gorilla/mux"
)

func main() {
	logFile, err := os.OpenFile("logFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE")
			w.Header().Add("Access-Control-Allow-Headers", "*")
		})
	r.HandleFunc("/items", Controllers.GetTodos).Methods("GET")
	r.HandleFunc("/items", Controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/items/{id}", Controllers.DeleteTodo).Methods("DELETE")

	fmt.Println("Server listening")
	http.ListenAndServe(":8081", r)
}
