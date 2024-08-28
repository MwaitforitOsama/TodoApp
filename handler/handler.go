package handler

import (
	"fmt"
	"net/http"
)

type Todo struct{}

func (o *Todo) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a Todo")
}

func (o *Todo) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Todo")
}

func (o *Todo) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get a Todo by ID")
}

func (o *Todo) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update a Todo by ID")
}

func (o *Todo) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a Todo by ID")
}
