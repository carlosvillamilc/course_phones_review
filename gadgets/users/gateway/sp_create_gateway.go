package gateway

import (
	"course-phones-review/gadgets/users/models"
	"course-phones-review/internal/database"
)

type UserCreateGateway interface {
	Create(cmd *models.CreateUserCMD) (*models.User, error)
	GetUserByID(userID int64) *models.User
	Authenticate(cmd *models.CreateUserCMD) (*models.User, error)
}

type UserCreateGtw struct {
	UserStorageGateway
}

func NewUserCreateGateway(client *database.MySqlClient) UserCreateGateway {
	return &UserCreateGtw{&UserStorage{client}}
}
