package model

import (
	"github.com/satori/uuid"
)

type Lecturer struct {
	LecturerID uuid.UUID
	Name,
	Gender,
	Address,
	DateOfBirth string
	Age     int32
	Subject Subject
}
