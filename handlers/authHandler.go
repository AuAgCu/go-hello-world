package handlers

import (
	"net/http"
	"strings"

	"app/service"

	"github.com/labstack/echo"
)

var AuthHandlerSingleton = &authHandlerImpl{}

type AuthHandler interface {
	Verify(c echo.Context) error
}

type authHandlerImpl struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandlerImpl{authService}
}

func (authHandler authHandlerImpl) Verify(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	jwtToken := strings.Replace(authHeader, "Bearer ", "", 1)

	return c.JSON(http.StatusOK, authHandler.authService.Verify(jwtToken))
}
