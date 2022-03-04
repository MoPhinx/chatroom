package process

import (
	"bufio"
	"fmt"
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
		fmt.Println("###############Welcom to Chat System##############")
		fmt.Println("\t\t\t1 Sign in")
		fmt.Println("\t\t\t2 Sign up")
		fmt.Println("\t\t\t3 Exit System")
		fmt.Println("\t\t\t Please choose(1 - 3): ")

		//接收用户选择
		_, err := fmt.Scanf("%d\n", &choose)
		if err != nil {
			fmt.Println("choose error, Please entry number")
			continue
		}
		//管理登录和注册
		up := &UserProcess{}

		switch choose {
		case 1: // 登录到聊天系统
			fmt.Println()
			fmt.Println("\t\t\t\tSign in to a chatroom")
			err := up.SignIn()
			if err != nil {
				return
			}
		case 2: //用户注册到聊天系统
			fmt.Println()
			fmt.Println("Sign up")
			err := up.SignUp()
			if err != nil {
				return
			}
		case 3: //退出聊天系统
			fmt.Println("Exit System")
			os.Exit(0)
		default: //输入有误，重新输入
			fmt.Println("You input is wrong, Please try again later!")
		}
	}
}

func (st *ShowTable) SignInMenu() {
	var key int

	var filePath string
	var smsProcess *SmSProcess

	for {
		fmt.Println("\t\t\t\tCongratulations on your successful login")
		fmt.Println("\t\t\t\t1, Displays an online list of users")
		fmt.Println("\t\t\t\t2, Send a message to Hall")
		fmt.Println("\t\t\t\t3, Send a message to Personal")
		fmt.Println("\t\t\t\t4, Message lists")
		fmt.Println("\t\t\t\t5, Back To Previous Menu")
		fmt.Println("\t\t\t\t6, Exit System")

		fmt.Println("Please entry (1 - 5):")
		_, err := fmt.Scanln(&key)
		if err != nil {
			return
		}
		switch key {
		case 1:
			//fmt.Println("Displays an online list of users")
			outputOnlineUser()
		case 2:
			//用输入缓冲区接收输入，不能直接用scanf或者scanln这类函数来接收,要么接收不全(空格接收不了)，要么
			//报错fmt.Scanln(&content) error= expected newline
			fmt.Println("Please entry content:")
			buff = bufio.NewReader(os.Stdin)

			content, err := buff.ReadString('\n')
			if err != nil {
				fmt.Println("buff.ReadString error=", err)
			}

			err = smsProcess.SendGroupMes(content)
			if err != nil {
				fmt.Println("smsProcess.SendGroupMes(content) error=", err)
				return
			}
		case 3:
			var userId int
			fmt.Println("Please entry Personal UserId:")
			_, err := fmt.Scanln(&userId)
			if err != nil {
				fmt.Println("fmt.Scanln(&userId) error=", err)
				return
			}
			fmt.Println("Please entry content:")
			buff = bufio.NewReader(os.Stdin)

			content, err := buff.ReadString('\n')
			if err != nil {
				fmt.Println("buff.ReadString error=", err)
			}
			err = smsProcess.SendPersonalMes(userId, content)
			if err != nil {
				fmt.Println("smsProcess.SendPersonalMes error=", err)
				return
			}

		case 4:
			fmt.Println("Message lists")
			mesMan := MessageMan{}
			fmt.Println("Please choose (1(Personal)-2(Group)):")
			var choose int
			_, err := fmt.Scanln(&choose)
			if err != nil {
				fmt.Println("fmt.Scanln(&choose) error=", err)
				return
			}
			if choose == 1 {
				filePath = "PersonalMes.txt"
			} else if choose == 2 {
				filePath = "GroupMes.txt"
			}
			mesMan.ReadMessageFromFile(filePath)
		case 5:
			fmt.Println("Back To Previous Menu")
			st := ShowTable{}
			st.MainInterface()
		case 6:
			fmt.Println("Exit System")
			os.Exit(0)
		default:
			fmt.Println("Your input error, Please entry right option")
		}
	}
}
