package interfaces

import "bd2-backend/src/models"

type IUser interface {
	GetUser(username string) models.User
	ValidateLogin() (bool, error)
	CreateUser() (int64, error)
}
