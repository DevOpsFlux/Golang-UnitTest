package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "testSwagger/docs"
)

// @title API Document
// @version 0.0.1
// @description this is sample document

func main() {
	e := echo.New()

	// localhost:1323/swagger/index.html
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
