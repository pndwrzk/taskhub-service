package error

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrPasswordIncorrect = errors.New("password incorrect")
var ErrUsernameExist = errors.New("username already exist")
var ErrEmailExist = errors.New("email already exist")
var ErrPasswordDoesNotMatch = errors.New("password does not match")
var ErrLogin = errors.New("email or password is incorrect")

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrEmailExist,
	ErrPasswordDoesNotMatch,
	ErrLogin,
}
