package registry

import (
	"database/sql"
	"my-github/clean-code-microservice-golang/adapter/controller"

	"github.com/go-redis/redis/v7"
)

type registry struct {
	db    *sql.DB
	redis *redis.Ring
}

type Registry interface {
	NewAppController() controller.AppController
}

type Option func(*registry)

func NewRedisClient(redis *redis.Ring) Option {
	return func(r *registry) {
		r.redis = redis
	}
}

func NewRegistry(db *sql.DB, options ...Option) Registry {
	regis := &registry{
		db: db,
	}

	for _, opt := range options {
		opt(regis)
	}
	return regis
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Student: r.NewStudentController(),
	}
}
