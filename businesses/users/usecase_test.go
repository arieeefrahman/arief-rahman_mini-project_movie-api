package users_test

import (
	"mini-project-movie-api/app/middlewares"
	"mini-project-movie-api/businesses/users"
	_userMock "mini-project-movie-api/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository	_userMock.Repository
	userService		users.UseCase

	userDomain		users.Domain
)

func TestMain(m *testing.M) {
	userService = users.NewUserUsecase(&userRepository, &middlewares.ConfigJWT{})

	userDomain = users.Domain{
		Email:    "test@test.com",
		Password: "123123",
	}

	m.Run()
}

func TestSignup(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		userRepository.On("Signup", &userDomain).Return(userDomain).Once()

		result := userService.Signup(&userDomain)

		assert.NotNil(t, result)
	})

	t.Run("Register | InValid", func(t *testing.T) {
		userRepository.On("Signup", &users.Domain{}).Return(users.Domain{}).Once()

		result := userService.Signup(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		userRepository.On("GetByEmail", &userDomain).Return(users.Domain{}).Once()

		result := userService.Login(&userDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		userRepository.On("GetByEmail", &users.Domain{}).Return(users.Domain{}).Once()

		result := userService.Login(&users.Domain{})

		assert.Empty(t, result)
	})
}