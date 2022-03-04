package message

const (
	//登录
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	//注册
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	//注销
	LogOffMesType    = "LogOffMes"
	LogOffResMesType = "LogOffResMes"
	//用户状态
	UserStateChangesMesType = "UserStateChangesMes"
	//聊天
	SmsMesType    = "SmsMes"    //群聊
	P2pSmsMesType = "P2pSmsMes" //点对点

)

//用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
	UserLearingIng
)

// Message 大的message种类
type Message struct {
	MType string `json:"m_type"` //message类型
	Data  string `json:"data"`   //message内容
}

// LoginMes 用于Login的message
type LoginMes struct {
	UserId    int    `json:"user_id"`    //用户ID
	UserPwd   string `json:"user_pwd"`   //用户密码
	UserName  string `json:"user_name"`  //用户名
	UserAddr  string `json:"user_addr"`  //用户地址
	UserPhone string `json:"user_phone"` //用户手机号
}

// LoginResMes 服务器端回复的message
type LoginResMes struct {
	Code  int   `json:"code"` //返回状态码，200表示登录成功，400表示用户未注册
	Users []int //返回用户Id的切片
	//UserName []string //返回用户名
	Error string `json:"error"` //返回错误信息,无错误则不返回
}

// RegisterMes 用于注册的message
type RegisterMes struct {
	User User `json:"user"`
}

// RegisterResMes 服务器端回复注册的Message
type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码：410表示已经占用；200表示注册成功;505表示未知错误
	Error string `json:"error"` // 返回错误信息
}

// LogOffMes  用于用户注销的message
type LogOffMes struct {
	User
}

// LogOffResMes server回复注销的Message
type LogOffResMes struct {
}

// UserStateChangesMes  配合服务器端推送用户状态变化的Message
type UserStateChangesMes struct {
	UserId int `json:"user_id"` // 用户Id
	Status int `json:"status"`  //用户状态
}

// SmsMes 群聊消息类型 //发送的message
type SmsMes struct {
	Content string `json:"content"` //发送的消息内容
	User           //匿名结构体，继承
}

//SmsResMes

// P2pSmsMes 点对点消息类型
type P2pSmsMes struct {
	UserIdByOther int //指定要发消息的对象的UserId
	SmsMes
}
