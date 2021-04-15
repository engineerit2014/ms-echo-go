package router

import (
	"github.com/labstack/echo/v4"
	middleKit "github.com/laironacosta/kit-go/middleware/echo"
	"github.com/laironacosta/ms-echo-go/controllers"
)

type Router struct {
	server          *echo.Echo
	userController  controllers.UserControllerInterface
	errorMiddleware middleKit.ErrorHandlerMiddlewareInterface
}

func NewRouter(
	server *echo.Echo,
	userController controllers.UserControllerInterface,
	errorMiddleware middleKit.ErrorHandlerMiddlewareInterface) *Router {
	return &Router{
		server,
		userController,
		errorMiddleware,
	}
}

func (r *Router) Init() {
	//create a default router with default middleware
	basePath := r.server.Group("/ms-echo-go")

	basePath.GET("/health", controllers.Health)

	users := basePath.Group("/users")
	{
		users.POST("", r.userController.Create, r.errorMiddleware.HandlerError)
		users.GET("/:email", r.userController.GetByEmail, r.errorMiddleware.HandlerError)
		users.PUT("/:email", r.userController.UpdateByEmail, r.errorMiddleware.HandlerError)
		users.PATCH("/:email", r.userController.UpdateByEmail, r.errorMiddleware.HandlerError)
		users.DELETE("/:email", r.userController.DeleteByEmail, r.errorMiddleware.HandlerError)
	}
}
