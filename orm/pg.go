package lib

/*
File name    : pg.go
Author       : miaoyc
Create Date  : 2023/4/30 23:48
Update Date  : 2023/4/30 23:48
Description  :
*/

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PGClient *gorm.DB
)

// 最大连接数
MAX_CONNECTION = 10000

// 最大空闲连接数
MAX_IDLE_CONNECTION = 1000

// InitPGConnection 初始化pg连接
func InitPGConnection(host, port, user, password string) {
	PGClient, _ = getPGClient(host, port, user, password)
	sqlDB, _ := PGClient.DB()
	sqlDB.SetMaxIdleConns(MAX_CONNECTION)
	sqlDB.SetMaxOpenConns(MAX_IDLE_CONNECTION)
}

// getPGClient 获取pg客户端
func getPGClient(host, port, user, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=tde sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 关闭gorm的默认打印日志
	// db.Logger = logger.Default.LogMode(logger.Silent)
	return db, nil
}
