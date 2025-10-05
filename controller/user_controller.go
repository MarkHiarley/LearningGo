package controller

import (
	"learninggo/models"
	"learninggo/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) userController {
	return userController{

		userUseCase: usecase,
	}
}

func (p *userController) CreateUser(ctx *gin.Context) {

	var user models.User
	err := ctx.BindJSON(&user)

	if user.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is obrigatory"})
		return
	}

	if user.Age == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "age is obrigatory"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err = p.userUseCase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)

}

func (p *userController) GetUsers(ctx *gin.Context) {
	users, err := p.userUseCase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
