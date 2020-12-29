package controller

import (
	"context"
	"my-github/clean-code-microservice-golang/domain/model"
	"my-github/clean-code-microservice-golang/usecase/student/interactor"
)

type StudentControllerImpl struct {
	studentInteractor interactor.StudentInteractor
}

type StudentController interface {
	CreateStudent(ctx context.Context, req *model.Student) error
}

func NewStudentController(std interactor.StudentInteractor) StudentController {
	return &StudentControllerImpl{studentInteractor: std}
}

func (sc *StudentControllerImpl) CreateStudent(ctx context.Context, req *model.Student) error {
	if err := sc.studentInteractor.CreateStudent(ctx, req); err != nil {
		return err
	}

	return nil

}
