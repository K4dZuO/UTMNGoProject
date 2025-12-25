package models

import (
	"github.com/google/uuid"
)

type ReviewInfo struct {
	ID    uuid.UUID
	ProductID int
	Rate int
}

