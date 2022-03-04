package process

//处理用户
import (
	"encoding/json"
	"fmt"
	"happiness999.cn/chatroom/server/model"
	"happiness999.cn/chatroom/server/utils"
	"happiness999.cn/chatroom/server/utils/message"
	"net"
)

// UserProcess 获取连接，处理登录和注册等逻辑
type UserProcess struct {
	Conn   net.Conn
	UserId int //表明该Coon是哪个用户的
}

// UserStateChanges 通知所有在线用户的func
func (p *UserProcess) UserStateChanges(userId int) {
	//userManage.UpdateUM()
	//遍历onlineusers，然后一个个发送UserStateChangesMes
	for id, up := range userManage.onlineUsersId {
		if id == userId {
			continue
		}
		up.UserState(userId)
	}
}

func (up *UserProcess) UserState(userId int) {

	//组装message
	var mes message.Message
	mes.MType = message.UserStateChangesMesType

	var userStateChangesMes message.UserStateChangesMes
	userStateChangesMes.UserId = userId
	userStateChangesMes.Status = message.UserOnline

	//将userStateChangesMes消息序列化存到mes.data里面
	data, err := json.Marshal(&userStateChangesMes)
	if err != nil {
		fmt.Println("json marshal error = ", err)
		return
	}
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(&mes)
	if err != nil {
		fmt.Println("json.Marshal error= ", err)
		return
	}

	//发送mes
	tf := utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("UserStateChanges error = ", err)
		return
	}
}

// ProcessLogin 处理用户登录逻辑
func (up *UserProcess) ProcessLogin(mes *message.Message) (err error) {
	//反序列化loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//声明一个用于服务器端回复的message
	var resMes message.Message
	resMes.MType = message.LoginResMesType

	var loginResMes message.LoginResMes

	//进行判断，如果验证成功则返回的message中带有200的状态码，否则带有500的状态码并给出error提示
	user, err := model.MyUserDao.SignIn(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ErrUserNotExists {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrUserPwd {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "server inside error"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user.UserName, "SignIn success")

		//将登录成功的用户的UserId赋值给 up
		up.UserId = loginMes.UserId
		//用户登录成功，把该登陆成功的用户放到UserManage的onlineUsers中
		userManage.AddOnlineUser(up)
		//通知其它用户，有人上线了，并发送更新的上线列表
		up.UserStateChanges(loginMes.UserId)
		//将登录成功的用户的Id放入到loginResMes.Users 切片中
		for id, _ := range userManage.onlineUsersId {
			loginResMes.Users = append(loginResMes.Users, id)
		}
	}

	//将loginResMes序列化并装到resMes中，再将resMes序列化并传输给client
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal error=", err)
		return
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal error=", err)
		return
	}

	//发送data给client
	tf := utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// ProcessReg 处理用户注册逻辑
func (up *UserProcess) ProcessReg(mes *message.Message) error {
	//1.先从mes中取出 mes.Data ，并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err := json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail error=", err)
		return err
	}

	//1先声明一个 resMes
	var resMes message.Message
	resMes.MType = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库去完成注册.
	//1.使用model.MyUserDao 到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ErrUserExists {
			registerResMes.Code = 410 //
			registerResMes.Error = model.ErrUserExists.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "An unknown error occurred in registration"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal error=", err)
		return err
	}

	//4. 将data 赋值给 resMes
	resMes.Data = string(data)

	//5. 对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal error=", err)
		return err
	}

	//6,发送data给client
	tf := utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)

	return nil
}
