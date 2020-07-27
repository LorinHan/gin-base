package serivces

import (
	"errors"
	"ginTest/models"
)

type userService struct{}

var UserService = &userService{}

func (us *userService) FindUserByPwd(password string) ([]*models.User, error) {
	users := models.FindUserByPwd(password)
	if len(users) <= 0 {
		return nil, errors.New("not found")
	}
	return users, nil
}
