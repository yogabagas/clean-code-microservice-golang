package interactor

import (
	"context"
	"my-github/clean-code-microservice-golang/domain/model"
	"my-github/clean-code-microservice-golang/internal/redis"
	"my-github/clean-code-microservice-golang/usecase/student/presenter"
	"my-github/clean-code-microservice-golang/usecase/student/repository"
)

type StudentInteractorImpl struct {
	studentPresenter  presenter.StudentPresenter
	studentCache      redis.InternalRedis
	studentRepository repository.StudentSQLRepository
}

type StudentInteractor interface {
	CreateStudent(ctx context.Context, req *model.Student) error
}

type Option func(std *StudentInteractorImpl)

func NewStudentRepositoryInteractor(sql repository.StudentSQLRepository) Option {
	return func(std *StudentInteractorImpl) {
		std.studentRepository = sql
	}
}

func NewStudentCacheInteractor(cache redis.InternalRedis) Option {
	return func(std *StudentInteractorImpl) {
		std.studentCache = cache
	}
}

func NewStudentInteractor(p presenter.StudentPresenter, options ...Option) StudentInteractor {
	interactor := &StudentInteractorImpl{
		studentPresenter: p,
	}

	for _, opt := range options {
		opt(interactor)
	}
	return interactor
}

func (s *StudentInteractorImpl) CreateStudent(ctx context.Context, req *model.Student) error {
	var err error

	if err = s.studentRepository.WriteStudent(ctx, req); err != nil {
		return err
	}

	if err = s.studentCache.Set(ctx, "student", req); err != nil {
		return err
	}

	return nil
}
