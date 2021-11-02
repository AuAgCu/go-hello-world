package handlers

import (
	"net/http"
	"strings"

	"app/service"

	"github.com/labstack/echo"
)

type AuthHandler interface {
	Verify(c echo.Context) error
}

type AuthHandlerImpl struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &AuthHandlerImpl{authService}
}

func (authHandler AuthHandlerImpl) Verify(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	jwtToken := strings.Replace(authHeader, "Bearer ", "", 1)

	return c.JSON(http.StatusOK, authHandler.authService.Verify(jwtToken))
}
