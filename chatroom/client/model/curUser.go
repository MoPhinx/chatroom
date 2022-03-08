package model

import (
	"happiness999.cn/chatroom/client/utils/message"
	"net"
)

// CurUser 维护当前用户的连接和用户信息
type CurUser struct {
	Conn net.Conn
	message.User
}
