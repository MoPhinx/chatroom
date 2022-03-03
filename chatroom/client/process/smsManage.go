package process

import (
	"encoding/json"
	"fmt"
	"happiness999.cn/chatroom/client/utils/message"
)

func outputGroupMes(mes *message.Message) { //处理sms
	//显示即可
	//1, 反序列化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal error=", err)
		return
	}

	//显示
	show := fmt.Sprintf("userID: %d, Mes: %s", smsMes.UserId, smsMes.Content)
	fmt.Println(show)
	fmt.Println()
}
