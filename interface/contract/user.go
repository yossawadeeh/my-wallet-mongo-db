package contract

import (
	"my-wallet-ntier-mongo/model"
	"my-wallet-ntier-mongo/response"
)

type UserRepository interface {
	GetUsers() (response []model.User, total int64, err error)
	GetUserById(userId string) (response *model.User, err error)
}

type UserService interface {
	GetUsers() (response []response.UserResponse, total int64, err error)
	GetUserById(userId string) (response *response.UserResponse, err error)
}
