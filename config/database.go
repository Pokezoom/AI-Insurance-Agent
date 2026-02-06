package config

import (
	"AI-Insurance-Agent/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return DB.AutoMigrate(&model.User{}, &model.AnalysisRecord{})
}
