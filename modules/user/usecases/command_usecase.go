package usecases

import (
	"errors"
	"log"
	"strings"

	"github.com/jeksilaen/api-builder/db"
	"github.com/jeksilaen/api-builder/modules/user/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserCommandUsecase struct {
	DB *gorm.DB
}

func NewUserCommandUsecase() *UserCommandUsecase {
	return &UserCommandUsecase{
		DB: db.GetDB(),
	}
}

func (uc *UserCommandUsecase) CreateUser(user *models.User) (*models.User, error) {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Create the user and handle duplicate error
	err = uc.DB.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("email or username already in use")
		}

		log.Println("Error creating user:", err)
		return nil, err
	}

	return user, nil
}

func (uc *UserCommandUsecase) FindUserByEmailAndPassword(req *models.LoginRequest) (*models.User, error) {
	// Find the user by email
	var user models.User
	err := uc.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}
