package service

import (
	"my-wallet-ntier-mongo/interface/contract"
	userResponse "my-wallet-ntier-mongo/response"
	"my-wallet-ntier-mongo/utils/mapper"
)

type userService struct {
	userRepo contract.UserRepository
}

func NewUserService(userRepo contract.UserRepository) contract.UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) GetUsers() (response []userResponse.UserResponse, total int64, err error) {
	users := []userResponse.UserResponse{}
	resp, total, err := u.userRepo.GetUsers()
	if err != nil {
		return nil, 0, err
	}

	for _, user := range resp {
		users = append(users, mapper.UserModelToResponse(user))
	}
	return users, total, err
}

func (u *userService) GetUserById(userId string) (response *userResponse.UserResponse, err error) {
	resp, err := u.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	user := mapper.UserModelToResponse(*resp)
	return &user, err
}
