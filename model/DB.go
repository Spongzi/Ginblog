package model

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db  *gorm.DB
	err error
)

// InitDb 配置数据库参数
func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DBUser,
		utils.DBPassword,
		utils.DBHost,
		utils.DBPort,
		utils.DBName)
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("数据库连接失败, err", err)
		return
	}
	err = db.AutoMigrate(&User{}, &Category{}, &Article{})
	if err != nil {
		fmt.Println("自动迁移失败", err)
		return
	}
	dbSet, err := db.DB()
	if err != nil {
		return
	}
	// 设置连接池中最大的闲置连接数
	dbSet.SetMaxIdleConns(10)
	// 设置数据库的最大连接数
	dbSet.SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	dbSet.SetConnMaxLifetime(10 * time.Second)
}
