package service

import (
	"testing"

	"my-wallet-ntier-mongo/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUsers() ([]model.User, int64, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]model.User), 0, args.Error(1)
	}

	return nil, 0, args.Error(1)
}
func (m *MockUserRepository) GetUserById(userId string) (*model.User, error) {
	args := m.Called(userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	if args.Get(0) != nil {
		return args.Get(0).(*model.User), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestGetUserById(t *testing.T) {
	t.Parallel()

	t.Run("GetUserById case success", func(t *testing.T) {

		mockUserRepo := new(MockUserRepository)
		mockUserService := NewUserService(mockUserRepo)

		userId := bson.NewObjectID()
		mockUser := &model.User{
			ID:        userId,
			FirstName: "John Doe",
			Email:     "john@example.com",
		}

		mockUserRepo.On("GetUserById", userId.Hex()).Return(mockUser, nil)

		// Act
		result, err := mockUserService.GetUserById(userId.Hex())

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, userId.Hex(), result.ID)
		assert.Equal(t, "John Doe", result.FirstName)
		assert.Equal(t, "john@example.com", result.Email)
		mockUserRepo.AssertExpectations(t)
	})
}
