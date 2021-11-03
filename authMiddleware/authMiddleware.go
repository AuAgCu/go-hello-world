package authmiddleware

import (
	"strings"

	"app/service"

	"github.com/labstack/echo"
)

var AuthMiddlewareEntity = &authMiddleWareImpl{service.AuthServiceSingleton}

type AuthMiddleWare interface {
	Verify(next echo.MiddlewareFunc)
}

type authMiddleWareImpl struct {
	authService service.AuthService
}

func (authMiddleWare authMiddleWareImpl) Verify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			jwtToken := strings.Replace(authHeader, "Bearer ", "", 1)
			authMiddleWare.authService.Verify(jwtToken)

			if err := next(c); err != nil {
				c.Error(err)
				return err
			}
			return nil
		}
	}
}
