package db

import (
	"fmt"

	"github.com/jeksilaen/api-builder/config"
	userModels "github.com/jeksilaen/api-builder/modules/user/models"	
	collectionModels"github.com/jeksilaen/api-builder/modules/collection/models"	
	requestModels"github.com/jeksilaen/api-builder/modules/request/models"	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	var err error
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrasi Model
	err = db.AutoMigrate(&userModels.User{}, &collectionModels.Collection{}, &requestModels.Request{})
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
