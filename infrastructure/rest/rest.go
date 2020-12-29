package rest

import (
	"database/sql"
	"fmt"
	"my-github/clean-code-microservice-golang/adapter/controller"
	"my-github/clean-code-microservice-golang/infrastructure/rest/group"
	"my-github/clean-code-microservice-golang/registry"
	"os"
	"os/signal"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/go-redis/redis/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type RestImpl struct {
	port          int
	db            *sql.DB
	router        *fiber.App
	appController *controller.AppController
}

func (r *RestImpl) GetAppController() *controller.AppController {
	return r.appController
}

func (r *RestImpl) GetRouter() *fiber.App {
	return r.router
}

func (r *RestImpl) Serve() {
	if err := r.router.Listen(fmt.Sprintf(":%d", r.port)); err != nil {
		panic(err)
	}
}

func NewRest(port int, db *sql.DB, rdb *redis.Ring) *RestImpl {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use("/swagger", swagger.Handler)
	app.Use(requestid.New())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		_ = app.Shutdown()
	}()

	registry := registry.NewRegistry(db, registry.NewRedisClient(rdb))
	appController := registry.NewAppController()

	r := &RestImpl{
		port:          port,
		db:            db,
		appController: &appController,
	}

	root := group.InitRoot(r)
	v1 := group.InitV1(r, root)
	group.InitStudentV1(r, v1)

	return r

}
