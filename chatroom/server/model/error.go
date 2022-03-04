package model

import "errors"

//自定义常见的错误
var (
	ErrUserNotExists = errors.New("\t\t\t\t\t\t 用户不存在，请先注册！")
	ErrUserExists    = errors.New("\t\t\t\t\t\t 用户已经存在！")
	ErrUserPwd       = errors.New("\t\t\t\t\t\t 密码错误！")
)
