package model

import "github.com/google/uuid"

type Student struct {
	StudentID uuid.UUID
	Name,
	Gender,
	Address,
	DateOfBirth string
	Age int32
}
