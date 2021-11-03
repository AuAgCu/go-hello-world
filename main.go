package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	firebaseUtil "app/firebase"
	"app/handlers"
	"app/repository"
	"app/service"
)

var tr = repository.NewUserReporitory(repository.Db)
var userService = service.NewUserService(tr)
var authService = service.NewAuthService()
var userHandler = handlers.NewUserHandler(userService)
var authHandler = handlers.NewAuthHandler(authService)

func main() {
	firebaseUtil.FirebaseUtil.InitFirebase()

	e := echo.New()
	e.Use(middleware.CORS())

	e.Static("/public/js", "public/js")

	e.File("/", "public/index.html")
	e.File("/login", "public/login.html")
	e.File("/signup", "public/sign-up.html")
	e.GET("/tasks", handlers.GetTasks)
	e.GET("/user/:id", userHandler.GetUser)

	e.POST("/user", userHandler.CreateUser)

	e.Start(":" + os.Getenv("PORT"))
	// e.Start(":8080")
}
