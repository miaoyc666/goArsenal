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
	"gorm.io/gorm/logger"
)

var (
	PGClient *gorm.DB
)

// MaxConnection 最大连接数
var MaxConnection = 10000

// MaxIdleConnection 最大空闲连接数
var MaxIdleConnection = 1000

// InitPGConnection 初始化pg连接
func InitPGConnection(host, port, user, password string) {
	PGClient, _ = getPGClient(host, port, user, password)
	sqlDB, _ := PGClient.DB()
	sqlDB.SetMaxIdleConns(MaxConnection)
	sqlDB.SetMaxOpenConns(MaxIdleConnection)
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
	db.Logger = logger.Default.LogMode(logger.Silent)
	return db, nil
}

// CountTable 表计数
func CountTable(tableName string) (int64, error) {
	var results []int64
	sql := fmt.Sprintf("SELECT count(*) as total FROM %s", tableName)
	db := lib.PGClient.Raw(sql).Scan(&results)
	if db.Error != nil {
		return 0, db.Error
	}
	return results[0], nil
}

