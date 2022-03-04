package process

import (
	"encoding/json"
	"fmt"
	"happiness999.cn/chatroom/client/utils/message"
)

//处理从server接收到的群聊消息
func outputGroupMes(mes *message.Message) error { //处理sms
	//显示即可
	//1, 反序列化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal error=", err)
		return err
	}

	//显示
	show := fmt.Sprintf("userID: %d, Mes: %s", smsMes.UserId, smsMes.Content)
	fmt.Println(show)
	fmt.Println()

	//存储
	storeMes := MessageMan{}
	filePath := "GroupMes.txt"
	err = storeMes.WriteMessageToFile(filePath, show)
	if err != nil {
		fmt.Println("WriteMessageToFile error=", err)
		return err
	}
	return nil
}

//处理从server接收到的个人消息
func outputPersonalMes(mes *message.Message) error {
	var p2p message.P2pSmsMes
	err := json.Unmarshal([]byte(mes.Data), &p2p)
	if err != nil {
		fmt.Println("json.Unmarshal error=", err)
		return err
	}

	//显示
	show := fmt.Sprintf("userID: %d, Mes: %s", p2p.UserId, p2p.Content)
	fmt.Println(show)
	fmt.Println()

	//存储
	storeMes := MessageMan{}
	filePath := "PersonalMes.txt"
	err = storeMes.WriteMessageToFile(filePath, show)
	if err != nil {
		fmt.Println("WriteMessageToFile error=", err)
		return err
	}
	return nil
}
