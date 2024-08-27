package model

import (
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
