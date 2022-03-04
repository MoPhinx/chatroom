package process

import (
	"encoding/json"
	"fmt"
	"happiness999.cn/chatroom/client/utils"
	"happiness999.cn/chatroom/client/utils/message"
	"net"
)

type ServerProcess struct {
	Conn net.Conn
}

// KeepConn 和服务器端保持连接
func (sp *ServerProcess) KeepConn() {
	tf := &utils.Transfer{
		Conn: sp.Conn,
	}
	for {
		//fmt.Println("client waiting for server message")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg error=", err)
			return
		}
		//如果读到新的message，则进行下一步处理
		switch mes.MType {
		case message.UserStateChangesMesType: //处理新用户上线
			//取出UserStateChangesMes.Data
			var userStateChangeMes message.UserStateChangesMes
			err := json.Unmarshal([]byte(mes.Data), &userStateChangeMes)
			if err != nil {
				fmt.Println("json.Unmarshal error=", err)
				return
			}
			//把这个用户的信息，状态保存到客户map[int]*User中
			updateUserStatus(&userStateChangeMes)
		case message.SmsMesType: //群发消息
			outputGroupMes(&mes)
		case message.P2pSmsMesType:
			err := outputPersonalMes(&mes)
			if err != nil {
				fmt.Println("outputPersonalMes error=", err)
				return
			}
		default:
			fmt.Println("Unknown message type")
		}
		fmt.Println("mes=", mes)
	}
}
