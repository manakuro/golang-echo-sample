package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/manakuro/golang-echo-sample/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", handler.Hello())
	e.GET("/users", handler.Users())

	log.Fatalln(e.Start(":8080"))
}
