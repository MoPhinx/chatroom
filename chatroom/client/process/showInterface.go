package process

import (
	"bufio"
	"fmt"
	"happiness999.cn/chatroom/client/utils/message"
	"os"
)

var buff *bufio.Reader

// ShowTable 用于显示主菜单和二级菜单
type ShowTable struct {
}

func (st *ShowTable) MainInterface() {
	var choose int // 接收用户选择
	for true {
		//打印主界面
		fmt.Println("####################################    欢迎来到happiness999.cn的原生聊天测试系统    ###################################")
		fmt.Println()
		fmt.Println("\t\t\t\t\t\t 1:注册")
		fmt.Println("\t\t\t\t\t\t 2:登录")
		fmt.Println("\t\t\t\t\t\t 3:退出")
		fmt.Printf("\t\t\t\t\t\t 请从(1-3)中选择：")

		//接收用户选择
		_, err := fmt.Scanf("%d\n", &choose)
		if err != nil {
			fmt.Println()
			fmt.Println()
			fmt.Println("\t\t\t\t\t\t 输入错误！请输入有效数字(1-3)")
			fmt.Println()
			continue
		}
		//管理登录和注册
		up := &UserProcess{}

		switch choose {
		case 1: //用户注册到聊天系统
			fmt.Println()
			fmt.Println()
			fmt.Println("\t\t\t\t\t\t 1,注册界面")
			err := up.SignUp()
			if err != nil {
				fmt.Println("SignUp error = ", err, "Please try again later")
				continue
				//return
			}
		case 2: // 登录到聊天系统
			fmt.Println()
			fmt.Println("\t\t\t\t\t\t 2,登录界面")
			err := up.SignIn()
			if err != nil {
				fmt.Println("SignIn error = ", err, "Please try again later")
				continue
				//return
			}
		case 3: //退出聊天系统
			fmt.Println("\t\t\t\t\t\t 您已退出系统")
			os.Exit(0)
		default: //输入有误，重新输入
			fmt.Println("\t\t\t\t\t\t 您的输入有误，请重新输入")
		}
	}
}

func (st *ShowTable) SignInMenu() {
	var key int

	var filePath string
	var smsProcess *SmSProcess

	for {
		fmt.Println("\t\t\t\t\t\t 1, 显示在线用户")
		fmt.Println("\t\t\t\t\t\t 2, 发送消息到大厅")
		fmt.Println("\t\t\t\t\t\t 3, 发送消息给个人")
		fmt.Println("\t\t\t\t\t\t 4, 显示消息列表")
		fmt.Println("\t\t\t\t\t\t 5, 退出登录")

		fmt.Println("\t\t\t\t\t\t 请从(1 - 5)中选择：")
		_, err := fmt.Scanln(&key)
		if err != nil {
			fmt.Println()
			fmt.Println()
			fmt.Println("\t\t\t\t\t\t 输入错误！请输入有效数字(1-5)")
			fmt.Println()
			continue
		}
		switch key {
		case 1: //显示在线用户
			fmt.Println()
			fmt.Println("\t\t\t\t\t\t 在线用户：")
			outputOnlineUser()
		case 2: //发送消息到大厅
			//用输入缓冲区接收输入，不能直接用scanf或者scanln这类函数来接收,要么接收不全(空格接收不了)，要么
			//报错fmt.Scanln(&content) error= expected newline
			fmt.Println()
			fmt.Printf("\t\t\t\t\t\t 请输入消息内容：")
			buff = bufio.NewReader(os.Stdin)

			content, err := buff.ReadString('\n')
			if err != nil {
				fmt.Print("buff.ReadString error=", err)
			}

			err = smsProcess.SendGroupMes(content)
			if err != nil {
				fmt.Println("smsProcess.SendGroupMes(content) error=", err)
				continue
			}
		case 3: //发送消息给个人
			var userId int
			fmt.Println()
			fmt.Printf("\t\t\t\t\t\t 请输入您想要发送用户的ID：")
			_, err := fmt.Scanln(&userId)
			if err != nil {
				fmt.Println("fmt.Scanln(&userId) error=", err, "Please try again!")
				continue
			}
			fmt.Printf("\t\t\t\t\t\t 请输入消息内容：")
			buff = bufio.NewReader(os.Stdin)

			content, err := buff.ReadString('\n')
			if err != nil {
				fmt.Println("buff.ReadString error=", err)
			}
			err = smsProcess.SendPersonalMes(userId, content)
			if err != nil {
				fmt.Println("smsProcess.SendPersonalMes error=", err)
				continue
			}

		case 4: //显示消息列表
			fmt.Println()
			fmt.Println("\t\t\t\t\t\t 消息列表")
			mesMan := MessageMan{}
			fmt.Printf("\t\t\t\t\t\t 请选择查看个人(1)消息还是大厅(2)消息：")
			var choose int
			_, err := fmt.Scanln(&choose)
			if err != nil {
				fmt.Println("fmt.Scanln(&choose) error=", err, "Please choose again")
				continue
			}
			if choose == 1 {
				filePath = "PersonalMes.txt"
			} else if choose == 2 {
				filePath = "GroupMes.txt"
			}
			err = mesMan.ReadMessageFromFile(filePath)
			if err != nil {
				fmt.Println("ReadMessageFromFile error = ", err)
			}
		case 5: //退出登录
			fmt.Println("\t\t\t\t\t\t 退出登录")
			fmt.Println()
			up := &UserProcess{}
			err := up.Logoff(CurUser.UserId)
			if err != nil {
				fmt.Println("Logoff error = ", err)
				continue
			}
			onlineUsers = make(map[int]*message.User, 10)
			st := ShowTable{}
			st.MainInterface()
		default:
			fmt.Println("\t\t\t\t\t\t 您的输入有误，请重新输入")
		}
	}
}
