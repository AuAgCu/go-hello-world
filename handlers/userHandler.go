package handlers

import (
	"net/http"
	"strconv"

	"app/models"
	"app/service"

	"github.com/labstack/echo"
)

type UserHandler interface {
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
}

type UserHandlerImpl struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{userService}
}

func (userHandler UserHandlerImpl) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	println(id)
	user := userHandler.userService.GetUserById(id)
	return c.JSON(http.StatusOK, user)
}

func (userHandler UserHandlerImpl) CreateUser(c echo.Context) error {
	param := new(models.User)
	//バインドしてJSON取得
	if err := c.Bind(param); err != nil {
		return err
	}

	println(param.FIRST_NAME)
	println(param.LAST_NAME)

	userHandler.userService.CreateUser(param.LAST_NAME, param.FIRST_NAME)

	return c.JSON(http.StatusOK, "hoge")
}
