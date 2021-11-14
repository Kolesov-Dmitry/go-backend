package groupstore

import (
	"time"

	"github.com/google/uuid"
)

type GroupDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
