package process

import (
	"encoding/json"
	"fmt"
	"happiness999.cn/chatroom/client/utils"
	"happiness999.cn/chatroom/client/utils/message"
	"net"
)

type UserProcess struct {
}

// SignIn 实现登录功能
func (up *UserProcess) SignIn() error {
	var userId int      //UserID
	var password string //User Password

	fmt.Printf("\t\t\t\t\t\t 请输入用户Id：")
	_, err2 := fmt.Scanln(&userId)
	if err2 != nil {
		fmt.Println("\t\t\t\t\t\t userId输入有误，请重新输入")
		return err2
	}
	fmt.Printf("\t\t\t\t\t\t 请输入用户密码：")
	_, err3 := fmt.Scanln(&password)
	if err3 != nil {
		fmt.Println("\t\t\t\t\t\t password输入有误，请重新输入")
		return err3
	}

	//连接到server
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("the net Dial error=", err)
		return err
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn Close error=", err)
		}
	}(conn)

	//通过conn发送自定义的message给server
	var mes message.Message          //自定义的message
	mes.MType = message.LoginMesType //设置message的类型

	var loginMes = &message.LoginMes{ //创建loginMes结构体，存放用于登录的用户信息
		UserId:  userId,
		UserPwd: password,
	}

	//序列化loginMes结构体 以便将其存放到 message.Data字段
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json Marshal error=", err)
		return err
	}
	mes.Data = string(data) //填充message.Data字段

	//序列化mes，以便通过Tcp/Ip传送到server
	data, err = json.Marshal(mes) //获取到要发送的数据的[]byte形式
	if err != nil {
		fmt.Println("json Marshal error=", err)
		return err
	}

	//将数据发送给server
	tfClient := &utils.Transfer{
		Conn: conn,
	}
	err = tfClient.WritePkg(data)
	if err != nil {
		fmt.Println("SingUp error=", err)
		return err
	}

	//处理server返回的数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Read Package error=", err)
		return err
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println()
		fmt.Println("\t\t\t\t\t\t 恭喜你，登录成功。")
		fmt.Println()
		//初始化CurUser
		CurUser.Conn = conn
		user := &message.User{
			UserId:     userId,
			UserStatus: message.UserOnline,
		}
		CurUser.User = *user

		//显示当前在线用户的列表
		fmt.Println("\t\t\t\t\t\t 当前在线用户列表：")
		for _, user := range loginResMes.Users {
			if user == userId {
				continue
			}
			fmt.Println("\t\t\t\t\t\t 当前在线用户ID：", user)

			//完成onlineUsers的初始化工作
			mUser := &message.User{
				UserId:     user,
				UserStatus: message.UserOnline,
			}
			onlineUsers[user] = mUser
		}

		fmt.Println()

		//启动一个协程用于与服务器端保持通讯,如果有数据推送给客户端则接收并显示
		sp := &ServerProcess{
			Conn: conn,
		}
		go func() {
			err := sp.KeepConn()
			if err != nil {
				fmt.Println("KeepConn error = ", err)
			}
		}()

		//显示菜单
		st := ShowTable{}
		st.SignInMenu()
	} else {
		fmt.Println(loginResMes.Error)
	}
	return nil
}

// SignUp 实现注册功能
func (up *UserProcess) SignUp() error {
	var id int
	var name string     //接收username
	var password string //接收password

	fmt.Printf("\t\t\t\t\t\t 请输入用户ID：") //获取用户Id
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("用户ID输入错误，请输入正确的用户ID，", err)
		return err
	}
	fmt.Printf("\t\t\t\t\t\t 请输入用户名：") //获取用户名
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Println("用户名输入错误，请输入正确的用户名， ", err)
		return err
	}
	fmt.Printf("\t\t\t\t\t\t 请输入用户密码：") //获取密码
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println("用户密码输入错误，请输入正确的用户密码， ", err)
		return err
	}

	//连接到server
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("the net Dial error=", err)
		return err
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("connect Close error= ", err)
		}
	}(conn)

	//通过conn发送自定义的message给server
	var mes message.Message             //自定义的message
	mes.MType = message.RegisterMesType //设置message的类型

	//创建RegisterMes结构体，存放用于登录的用户信息
	user := &message.User{
		UserId:   id,
		UserPwd:  password,
		UserName: name,
	}
	var registerMes = &message.RegisterMes{User: *user}

	//序列化registerMes结构体 以便将其存放到 message.Data字段
	data, err := json.Marshal(registerMes)

	if err != nil {
		fmt.Println("json Marshal error=", err)
		return err
	}

	mes.Data = string(data) //填充message.Data字段

	//序列化mes，以便通过Tcp/Ip传送到server
	data, err = json.Marshal(mes) //获取到要发送的数据的[]byte形式
	if err != nil {
		fmt.Println("json Marshal error=", err)
		return err
	}

	//将数据发送给server
	tfClient := &utils.Transfer{
		Conn: conn,
	}
	err = tfClient.WritePkg(data)
	if err != nil {
		fmt.Println("SingUp error=", err)
		return err
	}

	//接收服务器返回的message
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Read Package error=", err)
		return err
	}

	//接收服务器返回过来的message
	var registerRes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerRes)
	if registerRes.Code == 200 {
		fmt.Println("\t\t\t\t\t\t 注册成功！请登录")
		fmt.Println()
		fmt.Println()
		showTable := ShowTable{}
		showTable.MainInterface()
	} else {
		fmt.Println(registerRes.Error)
		showTable := ShowTable{}
		showTable.MainInterface()
	}
	return nil
}

// Logoff 实现注销功能
func (up *UserProcess) Logoff(userId int) error {
	//连接到server
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("the net Dial error=", err)
		return err
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn Close error=", err)
		}
	}(conn)

	var mes message.Message
	mes.MType = message.LogOffMesType

	//创建RegisterMes结构体，存放用于注销的用户信息
	user := &message.User{
		UserId: userId,
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(" json.Marshal(user) error = ", err)
		return err
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) error = ", err)
		return err
	}

	tf := utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg(data) error = ", err)
		return err
	}

	return nil
}
