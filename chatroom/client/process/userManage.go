package process

import (
	"fmt"
	"happiness999.cn/chatroom/client/model"
	"happiness999.cn/chatroom/client/utils/message"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //登录成功后完成对CurUser的初始化

//处理返回的UserStateChangesMes
func updateUserStatus(scm *message.UserStateChangesMes) {
	//如果map中没有这个user则创建并添加，如果有则不需要
	user, ok := onlineUsers[scm.UserId]
	if !ok {
		user = &message.User{UserId: scm.UserId}
	}
	user.UserStatus = scm.Status
	onlineUsers[scm.UserId] = user
	outputOnlineUser()
}

//显示当前在线用户
func outputOnlineUser() {
	//当前在线用户列表：
	fmt.Println("List of users")
	for id, _ := range onlineUsers {
		fmt.Println("user id:\t", id)
	}
}
