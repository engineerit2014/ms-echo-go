package router

import (
	"github.com/labstack/echo"
	"github.com/laironacosta/ms-echo-go/controllers"
)

type Router struct {
	server         *echo.Echo
	userController controllers.UserControllerInterface
}

func NewRouter(server *echo.Echo, userController controllers.UserControllerInterface) *Router {
	return &Router{
		server,
		userController,
	}
}

func (r *Router) Init() {
	//create a default router with default middleware
	basePath := r.server.Group("/ms-echo-go")

	basePath.GET("/health", controllers.Health)

	users := basePath.Group("/users")
	{
		users.POST("/", r.userController.Create)
		users.GET("/:email", r.userController.GetByEmail)
		users.PUT("/:email", r.userController.UpdateByEmail)
		users.PATCH("/:email", r.userController.UpdateByEmail)
		users.DELETE("/:email", r.userController.DeleteByEmail)
	}
}
