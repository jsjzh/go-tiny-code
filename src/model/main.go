package model

import (
	"fmt"
	"go-tiny-code/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func InitializeMysql(c *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Address,
		c.Mysql.Schema)

	if temp, error := gorm.Open(mysql.Open(dsn)); error != nil {
		panic(error.Error())
	} else {
		db = temp
		if sqlDB, error := db.DB(); error != nil {
			panic(error.Error())
		} else {
			// SetMaxIdleConns 设置空闲连接池中连接的最大数量
			sqlDB.SetMaxIdleConns(10)
			// SetMaxOpenConns 设置打开数据库连接的最大数量
			sqlDB.SetMaxOpenConns(100)
			// SetConnMaxLifetime 设置了连接可复用的最大时间
			sqlDB.SetConnMaxLifetime(time.Hour)
		}
	}

	return db
}
