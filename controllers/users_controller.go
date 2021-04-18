package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	respKit "github.com/laironacosta/kit-go/middleware/responses"
	"github.com/laironacosta/ms-echo-go/controllers/dto"
	"github.com/laironacosta/ms-echo-go/enums"
	"github.com/laironacosta/ms-echo-go/services"
	"github.com/laironacosta/ms-echo-go/translators"
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
	dictLang := c.Request().Context().Value(enums.I18nKey).(translators.DictLang)

	u := dto.CreateUserRequest{}
	if err := c.Bind(&u); err != nil {
		return respKit.GenericBadRequestError(enums.ErrorRequestBodyCode, err.Error())
	}

	if err := u.Validate(); err != nil {
		return respKit.GenericBadRequestError(enums.ErrorRequestBodyCode, err.Error())
	}

	if err := ctr.userService.Create(context.Background(), u); err != nil {
		return err
	}

	log.Infof("Request received: %+v \n", u)
	return c.JSON(http.StatusOK, dto.Response{
		Message: dictLang.GetMsg(enums.UserCreated),
	})
}

func (ctr *UserController) GetByEmail(c echo.Context) error {
	e := c.Param("email")
	fmt.Printf("Path param received: %+v \n", e)

	u, err := ctr.userService.GetByEmail(context.Background(), e)
	fmt.Printf("Service finished, controller\n")
	if err != nil {
		fmt.Printf("err %v\n", err)
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (ctr *UserController) UpdateByEmail(c echo.Context) error {
	dictLang := c.Request().Context().Value(enums.I18nKey).(translators.DictLang)

	u := dto.UpdateUserRequest{}
	if err := c.Bind(&u); err != nil {
		return respKit.GenericBadRequestError(enums.ErrorRequestBodyCode, err.Error())
	}

	if err := u.Validate(); err != nil {
		return respKit.GenericBadRequestError(enums.ErrorRequestBodyCode, err.Error())
	}

	e := c.Param("email")

	fmt.Printf("Request received: %+v \n", u)
	fmt.Printf("Path param received: %+v \n", e)

	if err := ctr.userService.UpdateByEmail(context.Background(), u, e); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: fmt.Sprintf(dictLang.GetMsg(enums.UserUpdated), e),
	})
}

func (ctr *UserController) DeleteByEmail(c echo.Context) error {
	dictLang := c.Request().Context().Value(enums.I18nKey).(translators.DictLang)

	e := c.Param("email")

	fmt.Printf("Path param received: %+v \n", e)

	if err := ctr.userService.DeleteByEmail(context.Background(), e); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: dictLang.GetMsg(enums.UserDeleted),
	})
}
