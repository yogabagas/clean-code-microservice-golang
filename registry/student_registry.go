package registry

import (
	"my-github/clean-code-microservice-golang/adapter/controller"
	ap "my-github/clean-code-microservice-golang/adapter/presenter"
	rsqlr "my-github/clean-code-microservice-golang/adapter/repository/sql"
	rds "my-github/clean-code-microservice-golang/internal/redis"
	"my-github/clean-code-microservice-golang/usecase/student/interactor"
	"my-github/clean-code-microservice-golang/usecase/student/presenter"
	"my-github/clean-code-microservice-golang/usecase/student/repository"
)

func (r *registry) NewStudentController() controller.StudentController {
	return controller.NewStudentController(
		r.NewStudentInteractor(),
	)
}

func (r *registry) NewStudentInteractor() interactor.StudentInteractor {
	return interactor.NewStudentInteractor(r.NewStudentPresenter(),
		interactor.NewStudentRepositoryInteractor(r.NewStudentSQLRepository()),
		interactor.NewStudentCacheInteractor(r.NewStudentCacheRepository()))
}

func (r *registry) NewStudentPresenter() presenter.StudentPresenter {
	return ap.NewUserPresenter()
}

func (r *registry) NewStudentSQLRepository() repository.StudentSQLRepository {
	return rsqlr.NewSQLRepository(r.db)
}

func (r *registry) NewStudentCacheRepository() rds.InternalRedis {
	return rds.NewRedisImpl(r.redis)
}
