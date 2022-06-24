package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	Id                   uuid.UUID
	AccountOriginId      string
	AccountDestinationId string
	Amount               float64
	Created_at           time.Time
}
