package models

import (
	"fmt"
	"gin-base/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func init() {
	dbConfig := conf.Database
	var err error
	db, err = gorm.Open(dbConfig.DbType, dbConfig.UserName+":"+dbConfig.Password+"@tcp("+dbConfig.Host+")/"+dbConfig.DbName+"?"+dbConfig.Args)
	if err != nil {
		fmt.Printf("mysql connect error %v\n", err)
		return
	}
	if db.Error != nil {
		fmt.Printf("database error %v\n", db.Error)
		return
	}
	db.LogMode(conf.Database.LogMode)
	db.DB().SetMaxIdleConns(dbConfig.MaxIdleConns) // 最大空闲连接数
	db.DB().SetMaxOpenConns(dbConfig.MaxOpenConns) // 最大连接数
	lifetime := dbConfig.MaxLifetime
	if lifetime == 0 {
		lifetime = 15 // 默认15秒超时
	}
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(lifetime)) // 建立连接的最大生命周期
}

func DBClose() {
	db.Close()
}

type tf func(db *gorm.DB) error

/**
 * @description: 传入匿名函数进行具体事务操作
 * @params: tf -> 具体操作
 * @return: error -> 不为空代表错误
 * @author: Lorin
 * @time: 2020/9/15 上午11:45
 */
func Tx(tf tf) error {
	tx := db.Begin()
	if err := tf(tx); err != nil {
		tx.Rollback()
		tx.Close()
		return err
	}
	tx.Commit()
	return nil
}
