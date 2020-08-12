package serivces

import (
	"fmt"
	"testing"
)

func Test_userService_FindUserByPwd(t *testing.T) {
	t.Run("FindUserByPwd", func(t *testing.T) {
		if users, err := FindUserByPwd("ff1234"); err != nil {
			fmt.Println(err)
		} else {
			for i := range users {
				fmt.Println(users[i])
			}
		}
	})
}
