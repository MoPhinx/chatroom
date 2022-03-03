package message

// User 定义用户信息结构体
type User struct {
	UserId     int    `json:"user_id"`
	UserPwd    string `json:"user_pwd"`
	UserName   string `json:"user_name"`
	UserStatus int    `json:"user_status"` //用户状态
	Sex        string `json:"sex"`
}
