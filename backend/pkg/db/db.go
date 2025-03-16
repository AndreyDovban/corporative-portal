package db

import (
	"backend/configs"
	"backend/pkg/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(sqlite.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		logger.ERROR("failed to connect database")
		panic("failed to connect database")
	}

	return &Db{db}
}
