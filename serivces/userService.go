package serivces

import (
	"errors"
	"gin-base/models"
)

type userService struct{}

//var UserService = &serService{}

func (userService) FindUserByPwd(password string) ([]*models.User, error) {
	users := models.FindUserByPwd(password)
	if len(users) <= 0 {
		return nil, errors.New("not found")
	}
	return users, nil
}
