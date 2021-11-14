package group

import (
	"github.com/google/uuid"
)

type Group struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
