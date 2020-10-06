package main

import (
	"handler"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middleware.Gzip())

	// Routes
	//e.GET("/", hello)

	// > http://localhost:1323/hello
	//e.GET("/hello", handler.MainPage())

	// > http://localhost:1323/hello/Hong-GilDong
	e.GET("/hello/:username", handler.MainPage2())

	/*
		// Routes
		v1 := e.Group("/api/v1")
		{
			v1.Post("/members", api.CreateMember)
			v1.Get("/members", api.GetMembers)
			v1.Get("/members/:id", api.GetMember)
		}
	*/

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler 1
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

/*
// Handler 2
func hello(c *echo.Context) error {
	var content struct {
		Response  string `json:"response"`
		Timestamp string `json:"timestamp"`
	}
	content.Response = "Hello, World!"
	content.Timestamp = getNow()
	return c.JSON(http.StatusOK, &content)
}
*/
