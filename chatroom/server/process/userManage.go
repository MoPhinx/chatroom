package process

import "fmt"

// UserManage 用来管理在线用户
type UserManage struct {
	onlineUsers map[int]*UserProcess
}

//UserManage 实例在服务器端，有且只有一个； 在多数地方都会使用到，我们将其定义为全局变量
var (
	userManage *UserManage
)

// 初始化userManage
func init() {
	userManage = &UserManage{onlineUsers: make(map[int]*UserProcess, 1024)}
}

// AddOnlineUser 对onlineUser添加和修改
func (um *UserManage) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

// DelOnlineUser 对OnlineUser的删除
func (um *UserManage) DelOnlineUser(up *UserProcess) {
	delete(um.onlineUsers, up.UserId)
}

// GetAllOnlineUsers 返回当前所有在线的用户
func (um *UserManage) GetAllOnlineUsers() map[int]*UserProcess {
	return um.onlineUsers
}

// GetOnlineUserById 根据ID返回对应的map的值
func (um *UserManage) GetOnlineUserById(userId int) (*UserProcess, error) {
	up, ok := um.onlineUsers[userId]
	if !ok {
		err := fmt.Errorf("the user %d don't exits", userId)
		return nil, err
	}
	return up, nil
}
