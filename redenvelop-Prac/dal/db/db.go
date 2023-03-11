package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var rdb *gorm.DB

var tableName = "rp_receive_record"

func InitDB() {
	dsn := "root:root2023@tcp(127.0.0.1:3306)/tech?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	rdb, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	sqlDB, err := rdb.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)
}
