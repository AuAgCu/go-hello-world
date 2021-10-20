package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"app/handlers"
	"app/repository"
	"app/service"
)

var tr = repository.NewUserReporitory(repository.Db)
var userService = service.NewUserService(tr)
var userHandler = handlers.NewUserHandler(userService)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks)
	e.GET("/user", userHandler.GetUser)

	e.POST("/user", userHandler.CreateUser)

	e.Start(":" + os.Getenv("PORT"))
}
