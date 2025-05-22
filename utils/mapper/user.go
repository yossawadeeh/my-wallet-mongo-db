package mapper

import (
	"my-wallet-ntier-mongo/model"
	"my-wallet-ntier-mongo/response"
)

func UserModelToResponse(user model.User) response.UserResponse {
	return response.UserResponse{
		ID:        user.ID.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address: response.AddressResponse{
			Line1:       user.Address.Line1,
			SubDistrict: user.Address.SubDistrict,
			District:    user.Address.District,
			Province:    user.Address.Province,
			Postcode:    user.Address.Postcode,
		},
	}
}
