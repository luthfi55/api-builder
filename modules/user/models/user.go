package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"type:uuid;primaryKey"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Username string `json:"username" gorm:"unique;not null" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Username string `json:"username" gorm:"unique;not null" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SucessRegistrationResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
	Links   []Link       `json:"links"`
}

type SucessLoginResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
	Token   string       `json:"token"`
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

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	return nil
}
