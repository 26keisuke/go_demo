package users

import (
	"strings"
	"github.com/26keisuke/go_demo/utils/errors"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return  errors.NewBadRequestError("invalid email address")
	}
	return nil
}
