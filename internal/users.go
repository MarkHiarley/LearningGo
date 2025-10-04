package users

import (
	"API/models"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []models.User

func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func CadastrarUser(ctx *gin.Context) {
	var body models.Body

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	novoUser := models.User{
		Id:   rand.Intn(101) + len(users),
		Name: body.Name,
		Age:  body.Age,
	}

	users = append(users, novoUser)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Usuário cadastrado com sucesso",
		"data":    novoUser,
	})
}
