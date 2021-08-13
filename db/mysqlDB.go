package db

import (
	"fmt"
	"gorm.io/driver/mysql" // mysql驱动
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func InitMysql() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库初始化失败！")
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("数据库连接池化失败！")
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}

func getMysqlDB() *gorm.DB {
	return db
}
