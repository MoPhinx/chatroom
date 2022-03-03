package process

import (
	"encoding/json"
	"fmt"
	"happiness999.cn/chatroom/client/utils"
	"happiness999.cn/chatroom/client/utils/message"
)

type SmSProcess struct {
}

// SendGroupMes 群发消息
func (p *SmSProcess) SendGroupMes(content string) error {
	//定义一个 Message
	var mes message.Message
	mes.MType = message.SmsMesType

	//定义一个 SmsMes
	var sms message.SmsMes
	sms.Content = content
	//定义SmsMes.User，并将其传给sms.User字段
	user := &message.User{
		UserId:     CurUser.UserId,
		UserStatus: CurUser.UserStatus,
	}
	sms.User = *user

	//序列化sms
	data, err := json.Marshal(sms)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal error=", err)
		return err
	}

	//将序列化后的sms传给mes.Data字段
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		return err
	}

	//将序列化后的mes发送给server
	tf := utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes WritePkg error=", err)
		return err
	}

	return nil
}
