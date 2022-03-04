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
	fmt.Println("\t\t\t\t\t\t 消息发送成功！")
	fmt.Println()

	//将自己发送的消息存储起来
	str := "me: " + content
	mesMan := MessageMan{}
	err = mesMan.WriteMessageToFile("GroupMes.txt", str)
	if err != nil {
		fmt.Println("WriteMessageToFile error=", err)
		return err
	}

	return nil
}

// SendPersonalMes 私聊
func (p *SmSProcess) SendPersonalMes(userIdByOther int, content string) error {
	//定义一个 Message
	var mes message.Message
	mes.MType = message.P2pSmsMesType

	//定义一个 P2pSmsMes
	var p2p message.P2pSmsMes
	p2p.UserIdByOther = userIdByOther

	//定义一个 P2pSmsMes.SmsMes
	var sms message.SmsMes
	sms.Content = content
	//定义 P2pSmsMes.SmsMes.User，并将其传给sms.User字段
	user := &message.User{
		UserId:     CurUser.UserId,
		UserStatus: CurUser.UserStatus,
	}
	sms.User = *user
	//将smsMes封装到p2pMes中
	p2p.SmsMes = sms

	//序列化p2p
	data, err := json.Marshal(p2p)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal error=", err)
		return err
	}

	//将序列化后的p2p传给mes.Data字段
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal error = ", err)
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

	fmt.Println("\t\t\t\t\t\t 消息已发送出去")
	fmt.Println()
	//将自己发送的消息存储起来
	str := "me: " + content
	mesMan := MessageMan{}
	err = mesMan.WriteMessageToFile("PersonalMes.txt", str)
	if err != nil {
		fmt.Println("WriteMessageToFile error=", err)
		return err
	}

	return nil
}
