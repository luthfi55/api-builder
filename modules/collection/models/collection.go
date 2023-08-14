package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/jeksilaen/api-builder/modules/user/models"
)

type Collection struct {
	gorm.Model
	ID       string `gorm:"type:uuid;primaryKey"`
	UserID   string `gorm:"type:uuid;not null"`
	Name     string `json:"name" validate:"required"`
	User     models.User   `gorm:"foreignKey:UserID" validate:"-"`
}

type CollectionResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Name    string `json:"name" validate:"required"`
}

type SucessCreateResponse struct {
	Message string       `json:"message"`
	Data    CollectionResponse `json:"data"`
	Links   []Link       `json:"links"`
}

type SucessGetResponse struct {
	Message string             `json:"message"`
	Data    []CollectionResponse `json:"data"`
	Links   []Link             `json:"links"`
}

type SucessDeleteResponse struct {
	Message string             `json:"message"`	
	Links   []Link             `json:"links"`
}

type FailedResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Links   []Link `json:"links"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func (collection *Collection) BeforeCreate(tx *gorm.DB) error {
	collection.ID = uuid.New().String()
	return nil
}