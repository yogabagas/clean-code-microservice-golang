package presenter

import (
	"context"
	"my-github/clean-code-microservice-golang/domain/model"
)

type StudentPresenter interface {
	ResponseStudent(ctx context.Context, resp *model.Student) (*model.Student, error)
}
