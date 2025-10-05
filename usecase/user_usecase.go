package usecase

import (
	"learninggo/models"
	"learninggo/services"
)

type UserUseCase struct {
	services services.UserServices
}

func NewUserUsecase(services services.UserServices) UserUseCase {
	return UserUseCase{
		services: services,
	}
}

func (pu *UserUseCase) GetUsers() ([]models.User, error) {
	return pu.services.GetUsers()
}

func (pu *UserUseCase) CreateUser(newUser models.User) (models.User, error) {

	userId, err := pu.services.CreateUser(newUser)

	if err != nil {
		return models.User{}, err
	}
	newUser.Id = userId

	return newUser, nil
}
