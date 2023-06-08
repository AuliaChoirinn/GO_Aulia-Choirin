package services

import (
	"GO_Aulia-Choirin/Praktikum/model"
	"GO_Aulia-Choirin/Praktikum/repositories"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserService(userRMock)
)

func TestGetUsersService_Success(t *testing.T) {
	usersMP := []*model.User{
		{
			Email:    "aulia.choirin.2003416@gmail.com",
			Password: "54321",
		},
		{
			Email:    "aulia.choirin.2003416@gmail.com",
			Password: "54321",
		},
	}

	usersM := []model.User{
		{
			Email:    "aulia.choirin.20034160@gmail.com",
			Password: "54321",
		},
		{
			Email:    "aulia.choirin.2003416@gmail.com",
			Password: "54321",
		},
	}

	userRMock.Mock.On("GetUsersRepository").Return(usersMP, nil)
	users, err := userSMock.GetUsersService()

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, usersM[0].Email, users[0].Email)
	assert.Equal(t, usersM[0].Password, users[0].Password)
}

func TestGetUsersService_Failure(t *testing.T) {
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserService(userRMock)
	userRMock.Mock.On("GetUsersRepository").Return(nil, errors.New("get all users failed"))
	users, err := userSMock.GetUsersService()

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestCreateUserService_Success(t *testing.T) {
	user := model.User{
		Email:    "aulia.choirin.2003416@gmail.com",
		Password: "54321",
	}

	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
	users, err := userSMock.CreateService(user)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Email, users.Email)
	assert.Equal(t, user.Password, users.Password)
}

func TestCreateUserService_Failure(t *testing.T) {
	user := model.User{
		Email:    "aulia.choirin.2003416@gmail.com",
		Password: "987654321",
	}

	userRMock.Mock.On("CreateRepository", user).Return(nil, fmt.Errorf("create user failed"))
	users, err := userSMock.CreateService(user)

	assert.Nil(t, users)
	assert.NotNil(t, err)
}
