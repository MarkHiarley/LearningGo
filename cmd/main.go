package main

import (
	"learninggo/controller"
	"learninggo/db"
	"learninggo/services"
	"learninggo/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	api := "/v1"

	dbconnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	UserServices := services.NewUseService(dbconnection)
	UserUseCase := usecase.NewUserUsecase(UserServices)
	UserController := controller.NewUserController(UserUseCase)

	server.GET(api+"/users", UserController.GetUsers)
	server.POST(api+"/users", UserController.CreateUser)
	server.Run(":8080")
}
