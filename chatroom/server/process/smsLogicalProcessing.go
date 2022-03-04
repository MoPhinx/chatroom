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

// ForwardMesToEverybody  转发消息给群聊
func (p *SmSProcess) ForwardMesToEverybody(mes *message.Message) error {
	//遍历服务器的onlineUsers map[int]*UserProcess，将消息一个个转发出去

	//取出Mes中的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		return err
	}

	//序列化mes
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("ForwardMes() json.Marshal error=", err)
		return err
	}

	for id, up := range userManage.onlineUsersId {
		//过滤掉自己，不要把再次消息发给自己
		if id == smsMes.UserId {
			continue
		}
		p.SendMesToEverybody(data, up.Conn)
	}

	return nil
}

// ForwardMesToOther 转发消息给个人
func (p *SmSProcess) ForwardMesToOther(mes *message.Message) error {
	var p2pMes message.P2pSmsMes
	err := json.Unmarshal([]byte(mes.Data), &p2pMes)
	if err != nil {
		return err
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("ForwardMes() json.Marshal error=", err)
		return err
	}

	for id, up := range userManage.onlineUsersId {
		if id == p2pMes.UserIdByOther {
			p.SendMesToEverybody(data, up.Conn)
		}
	}

	return nil
}

func (p *SmSProcess) SendMesToEverybody(mes []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(mes)
	if err != nil {
		fmt.Println("WritePkg Forward message error=", err)
		return
	}
}
