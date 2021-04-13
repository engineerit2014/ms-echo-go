package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/laironacosta/ms-echo-go/controllers/dto"
	"github.com/laironacosta/ms-echo-go/services"
	"net/http"
)

type UserControllerInterface interface {
	Create(c echo.Context) error
	GetByEmail(c echo.Context) error
	UpdateByEmail(c echo.Context) error
	DeleteByEmail(c echo.Context) error
}

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) UserControllerInterface {
	return &UserController{
		userService,
	}
}

func (ctr *UserController) Create(c echo.Context) error {
	u := dto.CreateUserRequest{}
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctr.userService.Create(context.Background(), u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("Request received: %+v \n", u)
	return c.JSON(http.StatusOK, dto.Response{
		Message: "created",
	})
}

func (ctr *UserController) GetByEmail(c echo.Context) error {
	e := c.Param("email")
	fmt.Printf("Path param received: %+v \n", e)

	u, err := ctr.userService.GetByEmail(context.Background(), e)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func (ctr *UserController) UpdateByEmail(c echo.Context) error {
	u := dto.UpdateUserRequest{}
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	e := c.Param("email")

	fmt.Printf("Request received: %+v \n", u)
	fmt.Printf("Path param received: %+v \n", e)

	if err := ctr.userService.UpdateByEmail(context.Background(), u, e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, dto.Response{
		Message: "updated",
	})
}

func (ctr *UserController) DeleteByEmail(c echo.Context) error {
	e := c.Param("email")

	fmt.Printf("Path param received: %+v \n", e)

	if err := ctr.userService.DeleteByEmail(context.Background(), e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, dto.Response{
		Message: "deleted",
	})
}
