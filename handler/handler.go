package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Users() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := User{
			Name:  "myname",
			Email: "email 2",
		}

		return c.JSON(http.StatusOK, u)
	}
}
