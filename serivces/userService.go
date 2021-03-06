package serivces

import (
	"errors"
	"gin-base/models"
)

func FindUserByPwd(password string) ([]*models.User, error) {
	users := models.FindUserByPwd(password)
	if len(users) <= 0 {
		return nil, errors.New("not found")
	}
	return users, nil
}
