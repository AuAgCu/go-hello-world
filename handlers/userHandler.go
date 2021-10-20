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
	userHandler.userService.GetUserById(id)
	return c.JSON(http.StatusOK, models.GetUser())
}

func (userHandler UserHandlerImpl) CreateUser(c echo.Context) error {
	println("hogehoge")
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")

	println(firstName)
	println(lastName)

	userHandler.userService.CreateUser(lastName, firstName)

	return c.JSON(http.StatusOK, "hoge")
}
