package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:8099") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// # run example.go and visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser
// #$ go run test.go
