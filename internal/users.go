package users

import (
	"API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {

	data := []models.User{
		{
			Id:   1,
			Name: "teste,",
			Age:  12,
		},
		{
			Id:   1,
			Name: "teste,",
			Age:  12,
		},
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
