package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	_ "github.com/[각자프로젝트_위치]/docs"
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