package services

import "github.com/AkifhanIlgaz/grpc-server/models"

type AuthService interface {
	SignUp(*models.SignUpInput) (*models.DBResponse, error)
	SignIn(*models.SignInInput) (*models.DBResponse, error)
}
