package service

import (
	"back-end/conf"
	"back-end/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	dbConf := conf.Config.DB
	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/eth_wallet?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("opens database failed: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("get sql DB failed: " + err.Error())
	}
	// 自动迁移(创建表)
	// https://gorm.io/zh_CN/docs/migration.html
	if err = db.AutoMigrate(model.Models...); err != nil {
		log.Fatalf("auto migrate failed" + err.Error())
	}
	sqlDB.SetMaxIdleConns(dbConf.MaxIdle)                                     // 空闲连接池连接最大数量
	sqlDB.SetMaxOpenConns(dbConf.MaxActive)                                   // 打开数据库链接最大数量
	sqlDB.SetConnMaxLifetime(time.Duration(dbConf.MaxLifeTime) * time.Second) // 可复用连接最大时间
	log.Print("连接数据库成功")
}
