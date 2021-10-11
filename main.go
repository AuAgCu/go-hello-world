package main

import (
	"os"

	"github.com/labstack/echo"

	"app/handlers"
)

func main() {
	e := echo.New()
	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks)
	e.GET("/user", handlers.GetUsers)

	e.Start(":" + os.Getenv("PORT"))
}
