package process

import (
	"encoding/json"
	"fmt"
	"happiness999.cn/chatroom/server/utils"
	"happiness999.cn/chatroom/server/utils/message"
	"net"
)

type SmSProcess struct {
}

// ForwardMes 转发消息
func (p *SmSProcess) ForwardMes(mes *message.Message) {
	//遍历服务器的onlineUsers map[int]*UserProcess，将消息一个个转发出去

	//取出Mes中的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return
	}

	//序列化mes，将其
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("ForwardMes() json.Marshal error=", err)
		return
	}

	for id, up := range userManage.onlineUsers {
		//过滤掉自己，不要把再次消息发给自己
		if id == smsMes.UserId {
			continue
		}
		p.SendMesToEverybody(data, up.Conn)
	}
}

func (p SmSProcess) SendMesToEverybody(mes []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(mes)
	if err != nil {
		fmt.Println("WritePkg Forward message error=", err)
		return
	}
}
