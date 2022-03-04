package process

import (
	"fmt"
	"happiness999.cn/chatroom/server/utils"
	"happiness999.cn/chatroom/server/utils/message"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// Process 不断读取客户端发来的TCP字节流
func (p *Processor) Process() {
	defer func(Conn net.Conn) {
		err := Conn.Close()
		if err != nil {

		}
	}(p.Conn)

	//循环的读客户端发送的信息
	for {
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("readPkg error=", err)
			return
		}

		//fmt.Println("message=", mes)
		p := &Processor{Conn: p.Conn}
		err = p.KindOfMes(&mes)
		if err != nil {
			fmt.Println("KindOfMes error=", err)
			return
		}
	}
}

// KindOfMes 根据不同的message种类，选择调用不同的function,处理不同的逻辑
func (p *Processor) KindOfMes(mes *message.Message) (err error) {
	fmt.Println("mes=", mes)
	switch mes.MType {
	case message.LoginMesType:
		//处理登录逻辑
		up := &UserProcess{
			Conn: p.Conn,
		}
		err = up.ProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册逻辑
		up := &UserProcess{
			Conn: p.Conn,
		}
		err = up.ProcessReg(mes)
	case message.SmsMesType:
		//创建一个SmsProcess的实例，完成转发群聊消息的任务
		smSProcess := SmSProcess{}
		err := smSProcess.ForwardMesToEverybody(mes)
		if err != nil {
			return err
		}
	case message.P2pSmsMesType:
		smSProcess := SmSProcess{}
		err := smSProcess.ForwardMesToOther(mes)
		if err != nil {
			return err
		}
	default:
		fmt.Println("the kind of message don't exits, can't handle it")
	}
	return
}
