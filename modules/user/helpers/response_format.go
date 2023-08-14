package helpers

import (
	"github.com/jeksilaen/api-builder/middlewares"
	"github.com/jeksilaen/api-builder/modules/user/models"
)

func ReturnFailedRegisterResponse(message string) *models.FailedResponse {
	return &models.FailedResponse{
		Error:   "Register failed",
		Message: message,
		Links: []models.Link{
			{
				Rel:  "login",
				Href: "/users/v1/login",
			},
		},
	}
}

func ReturnFailedLoginResponse(message string) *models.FailedResponse {
	return &models.FailedResponse{
		Error:   "Login failed",
		Message: message,
		Links: []models.Link{
			{
				Rel:  "register",
				Href: "/users/v1/register",
			},
		},
	}
}

func ReturnSucessRegisterResponse(createdUser *models.User) *models.SucessRegistrationResponse {
	return &models.SucessRegistrationResponse{
		Message: "Registered user sucessfully",
		Data: models.UserResponse{
			ID:       createdUser.ID,
			Email:    createdUser.Email,
			Username: createdUser.Username,
		},
		Links: []models.Link{
			{
				Rel:  "login",
				Href: "/users/v1/login",
			},
		},
	}
}

func ReturnSucessLoginResponse(user *models.User) *models.SucessLoginResponse {
	token, err := middlewares.GenerateToken(user.ID)
	if err != nil {
		return &models.SucessLoginResponse{
			Message: "Login user sucessfully",
			Data: models.UserResponse{
				ID:       user.ID,
				Email:    user.Email,
				Username: user.Username,
			},
			Token: "Failed to generate token, please try again",
		}
	}

	return &models.SucessLoginResponse{
		Message: "Login user sucessfully",
		Data: models.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
		Token: token,
	}
}
