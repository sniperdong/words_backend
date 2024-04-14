package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	dsn     = "root:PSxf@520@tcp(localhost:3306)/words?charset=utf8&parseTime=True&loc=Local"
	newsDSN = "root:PSxf@520@tcp(localhost:3306)/news?charset=utf8&parseTime=True&loc=Local"
)

var (
	db     *gorm.DB
	newsDB *gorm.DB
)

func Init() {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	newsDB, err = gorm.Open(mysql.Open(newsDSN), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
