package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	// Static Route
	server.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	server.POST("/login", func(context *gin.Context) {
		context.String(http.StatusOK, "hello login")
	})

	// Parameter in the Path Route
	server.GET("/users/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "This is passed name: %s", name)
	})

	// Query value
	server.GET("/order", func(context *gin.Context) {
		id := context.Query("id")
		context.String(http.StatusOK, "Order # is %s", id)
	})

	server.GET("/views/*.html", func(context *gin.Context) {
		path := context.Param(".html")
		context.String(http.StatusOK, "Matched value is %s", path)
	})

	// Will learn later, it is used for setup different server; They are independent.

	// The Execution Order really matter, this block has to be executed before server.Run(":8080"),
	// Otherwise, the new port for 8081 will not execute. Not Quite Understand this part yet.
	go func() {
		server1 := gin.Default()
		server1.GET("/hello1", func(context *gin.Context) {
			context.String(http.StatusOK, "hello world 1 ")
		})
		server1.Run(":8081")
	}()

	// if it does not pass any value, it will consume the default 8080
	server.Run(":8080")
	// Caution: Please do not forget to add the ":" for port num
	// Incorrect: server.Run("8080")

}
