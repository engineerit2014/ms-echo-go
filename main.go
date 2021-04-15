package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	middleKit "github.com/laironacosta/kit-go/middleware/echo"
	pgKit "github.com/laironacosta/kit-go/postgresql"
	"github.com/laironacosta/ms-echo-go/controllers"
	"github.com/laironacosta/ms-echo-go/migrations"
	repo "github.com/laironacosta/ms-echo-go/repository"
	"github.com/laironacosta/ms-echo-go/router"
	"github.com/laironacosta/ms-echo-go/services"
	"github.com/pkg/errors"
)

// cfg is the struct type that contains fields that stores the necessary configuration
// gathered from the environment.
var cfg struct {
	DBUser string `envconfig:"DB_USER" default:"root"`
	DBPass string `envconfig:"DB_PASS" default:"root"`
	DBName string `envconfig:"DB_NAME" default:"user"`
	DBHost string `envconfig:"DB_HOST" default:"db"`
	DBPort int    `envconfig:"DB_PORT" default:"5432"`
}

func main() {
	echo := echo.New()
	echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human={latency_human}\n",
	}))
	echo.Use(middleware.Recover())

	if err := envconfig.Process("LIST", &cfg); err != nil {
		err = errors.Wrap(err, "parse environment variables")
		return
	}

	db := pgKit.NewPgDB(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
		User:     cfg.DBUser,
		Password: cfg.DBPass,
		Database: cfg.DBName,
	})
	migrations.Init(db)

	userRepo := repo.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	m := middleKit.NewErrorHandlerMiddleware()

	r := router.NewRouter(echo, userController, m)
	r.Init()

	echo.Start(":8080") // listen and serve on 0.0.0.0:8080
}
