package process

import (
	"fmt"
	"os"
)

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
			return
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
	var content string

	var smsProcess *SmSProcess

	for {
		fmt.Println("\t\t\t\tCongratulations on your successful login")
		fmt.Println("\t\t\t\t1, Displays an online list of users")
		fmt.Println("\t\t\t\t2, Send a message")
		fmt.Println("\t\t\t\t3, Message lists")
		fmt.Println("\t\t\t\t4, Back To Previous Menu")
		fmt.Println("\t\t\t\t5, Exit System")

		fmt.Println("Please entry (1 - 5):")
		fmt.Scanln(&key)
		switch key {
		case 1:
			//fmt.Println("Displays an online list of users")
			outputOnlineUser()
		case 2:
			fmt.Println("Please entry content:")
			_, err := fmt.Scanln(&content)
			if err != nil {
				return
			}
			err = smsProcess.SendGroupMes(content)
			if err != nil {
				return
			}

		case 3:
			fmt.Println("Message lists")
		case 4:
			fmt.Println("Back To Previous Menu")
			st := ShowTable{}
			st.MainInterface()
		case 5:
			fmt.Println("Exit System")
			os.Exit(0)

		default:
			fmt.Println("Your input error, Please entry right option")
		}
	}
}
