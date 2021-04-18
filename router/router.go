package router

import (
	"github.com/labstack/echo/v4"
	middleKit "github.com/laironacosta/kit-go/middleware/echo"
	"github.com/laironacosta/ms-echo-go/controllers"
	"github.com/laironacosta/ms-echo-go/middlewares"
)

type Router struct {
	server          *echo.Echo
	userController  controllers.UserControllerInterface
	errorMiddleware middleKit.ErrorHandlerMiddlewareInterface
	i18nMiddleware  middlewares.I18nMiddlewareInterface
}

func NewRouter(
	server *echo.Echo,
	userController controllers.UserControllerInterface,
	errorMiddleware middleKit.ErrorHandlerMiddlewareInterface,
	i18nMiddleware middlewares.I18nMiddlewareInterface,
) *Router {
	return &Router{
		server,
		userController,
		errorMiddleware,
		i18nMiddleware,
	}
}

func (r *Router) Init() {
	//create a default router with default middlewares
	basePath := r.server.Group("/ms-echo-go")

	basePath.GET("/health", controllers.Health)

	users := basePath.Group("/users", r.errorMiddleware.HandlerError, r.i18nMiddleware.HandlerError)
	{
		users.POST("", r.userController.Create)
		users.GET("/:email", r.userController.GetByEmail)
		users.PUT("/:email", r.userController.UpdateByEmail)
		users.PATCH("/:email", r.userController.UpdateByEmail)
		users.DELETE("/:email", r.userController.DeleteByEmail)
	}
}
