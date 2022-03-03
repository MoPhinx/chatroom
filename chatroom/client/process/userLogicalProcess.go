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

	fmt.Printf("Please entry your userId:")
	_, err2 := fmt.Scanln(&userId)
	if err2 != nil {
		return err2
	}
	fmt.Printf("Please entry your password:")
	_, err3 := fmt.Scanln(&password)
	if err3 != nil {
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

		//初始化CurUser
		CurUser.Conn = conn
		user := &message.User{
			UserId:     userId,
			UserStatus: message.UserOnline,
		}
		CurUser.User = *user

		//显示当前在线用户的列表
		fmt.Println("The current online list of users is as follows:")
		for _, user := range loginResMes.Users {
			if user == userId {
				continue
			}
			fmt.Println("the user id = ", user)

			//完成onlineUsers的初始化工作
			mUser := &message.User{
				UserId:     user,
				UserStatus: message.UserOnline,
			}
			onlineUsers[user] = mUser
		}

		//启动一个协程用于与服务器端保持通讯,如果有数据推送给客户端则接收并显示
		sp := &ServerProcess{
			Conn: conn,
		}
		go sp.KeepConn()

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

	fmt.Println("Please entry User Id") //获取用户Id
	_, err := fmt.Scanln(&id)
	if err != nil {
		return err
	}
	fmt.Println("Please entry User Name") //获取用户名
	_, err = fmt.Scanln(&name)
	if err != nil {
		return err
	}
	fmt.Println("Please entry User Password") //获取密码
	_, err = fmt.Scanln(&password)
	if err != nil {
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
		fmt.Println("Sign Up success, Please Sign in again")
		showTable := ShowTable{}
		showTable.MainInterface()
	} else {
		fmt.Println(registerRes.Error)
		showTable := ShowTable{}
		showTable.MainInterface()
	}
	return nil
}
