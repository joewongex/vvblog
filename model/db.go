package model

import (
	"fmt"
	"vvblog/config"
	"vvblog/vlog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	cfg := config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		vlog.Fatalf("[数据库]连接数据库失败：%v", err)
	}

	if err = DB.AutoMigrate(&User{}, &Post{}, &PostCategory{}); err != nil {
		vlog.Fatalf("[数据库]数据库自动迁移失败：%v", err)
	}
}
