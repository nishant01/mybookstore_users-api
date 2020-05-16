package services

import (
	"github.com/nishant01/mybookstore_users-api/doamin/users"
	"github.com/nishant01/mybookstore_users-api/utils/errors"
)

func GetUserById(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
