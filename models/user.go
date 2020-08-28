package models

import (
	"log"
)

type User struct {
	ID       int64
	Username string `gorm:"column:username" form:"username" json:"username" binding:"required"`
	Password string `gorm:"column:password" form:"password" json:"password" binding:"required"`
}

func FindUserByPwd(password string) []*User {
	var users []*User
	db.Where("password = ?", password).Offset(1).Limit(2).Find(&users)
	return users
}

func FindUserById(id int64) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUsersNative(begin int, size int) []*User {
	rows, err := db.Raw("select id, username, password from users limit ?, ?", begin, size).Rows()
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	users := make([]*User, 0, 10)
	for rows.Next() {
		var id int64
		var name, pwd string
		rows.Scan(&id, &name, &pwd)
		user := &User{id, name, pwd}
		users = append(users, user)
	}
	return users
}
func TestRows2Maps(begin int, size int) *[]*map[string]string {
	rows, err := db.Raw("select * from users limit ?, ?", begin, size).Rows()
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	result := RowsToMaps(rows)
	return result
}

