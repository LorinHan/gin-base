package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init()  {
	//dbConfig := conf.Conf.Database
	//var err error
	//DB, err = gorm.Open(dbConfig.DbType, dbConfig.UserName + ":" + dbConfig.Password +"@/" + dbConfig.DbName + "?" + dbConfig.Args)
	//if err != nil {
	//	fmt.Printf("mysql connect error %v", err)
	//	return
	//}
	//if DB.Error != nil {
	//	fmt.Printf("database error %v", DB.Error)
	//	return
	//}

	//DB.LogMode(true)
	//DB.DB().SetMaxIdleConns(10) // 最大空闲连接数
	//DB.DB().SetMaxOpenConns(50) // 最大连接数
}
