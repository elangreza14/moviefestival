package dto

import "github.com/google/uuid"

type UpdateOrder struct {
	OrderID uuid.UUID
	UserID  uuid.UUID
	CartIDs []int
}
