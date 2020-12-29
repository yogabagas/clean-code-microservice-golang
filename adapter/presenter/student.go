package presenter

import (
	"context"
	"my-github/clean-code-microservice-golang/domain/model"
	"my-github/clean-code-microservice-golang/usecase/student/presenter"
)

type StudentPresenterImpl struct{}

func NewUserPresenter() presenter.StudentPresenter {
	return &StudentPresenterImpl{}
}

func (s *StudentPresenterImpl) ResponseStudent(ctx context.Context, resp *model.Student) (*model.Student, error) {
	return resp, nil
}
