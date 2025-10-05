package main

import (
	users "API/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	api := "/v1"

	server.GET(api+"/users", users.GetUsers)
	server.POST(api+"/users", users.CadastrarUser)

	server.Run(":8080")
}
