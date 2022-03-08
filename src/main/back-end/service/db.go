package service

import (
	"back-end/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := "root:root@tcp(127.0.0.1:3306)/eth_wallet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("opens database failed: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("get sql DB failed: " + err.Error())
	}
	// 自动创建数据库
	// https://gorm.io/zh_CN/docs/migration.html
	if err = db.AutoMigrate(model.Models...); err != nil {
		log.Fatalf("auto migrate failed" + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)                  // 空闲连接池连接最大数量
	sqlDB.SetMaxOpenConns(50)                  // 打开数据库链接最大数量
	sqlDB.SetConnMaxLifetime(30 * time.Second) // 可复用连接最大事件
	log.Print("连接数据库成功")
}
