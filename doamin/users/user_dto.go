package users

import (
	"strings"

	"github.com/nishant01/mybookstore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

// User decleared
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address.")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password address.")
	}
	return nil
}
