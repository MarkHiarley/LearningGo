package main

import (
	users "API/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/users", users.GetUsers)

	server.Run(":8080")
}
