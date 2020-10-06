package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func MainPage2() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username") //플레이스홀더 username의 값 추출
		return c.String(http.StatusOK, "Hello World "+username)
	}
}
