package repository

import (
	"context"
	"my-github/clean-code-microservice-golang/domain/model"
)

type StudentSQLRepository interface {
	WriteStudent(ctx context.Context, req *model.Student) error
}

type StudentCacheRepository interface {
	WriteStudent(ctx context.Context, key string, value interface{}) error
}
