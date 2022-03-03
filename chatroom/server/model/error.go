package model

import "errors"

//自定义常见的错误
var (
	ErrUserNotExists = errors.New("user don't exists")
	ErrUserExists    = errors.New("user already exists")
	ErrUserPwd       = errors.New("password wrong")
)
