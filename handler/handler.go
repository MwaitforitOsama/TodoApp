package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id          uuid.UUID  `json:"id"`
	Status      string     `json:"first_name"`
	Description string     `json:"last_name"`
	Title       string     `json:"-"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

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
