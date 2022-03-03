package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"happiness999.cn/chatroom/server/utils/message"
)

var (
	MyUserDao *UserDao
)

// UserDao 完成对User的各种操作
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 使用工厂模式，创建UserDao实例
func NewUserDao(pool *redis.Pool) *UserDao {
	userDao := &UserDao{pool: pool}
	return userDao
}

//根据一个UserId返回一个User实例和error
func (d *UserDao) getUserById(conn redis.Conn, userId int) (*message.User, error) {
	res, err := redis.String(conn.Do("HGet", "users", userId))

	if err != nil {
		if err == redis.ErrNil {
			err = ErrUserNotExists
		}
		return nil, err
	}

	//将res反序列化为User实例
	user := &message.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return nil, err
	}

	return user, nil
}

// SignIn 完成登录校验；若成功则返回一个User实例，若id或Pwd有错误则返回对应的错误信息
func (d *UserDao) SignIn(userid int, userPwd string) (*message.User, error) {
	//从连接池中取出一个连接
	conn := d.pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("conn.Close err = ", err)
		}
	}(conn)

	user, err := d.getUserById(conn, userid) //获取用户
	if err != nil {
		return nil, err
	}

	if user.UserPwd != userPwd {
		err = ErrUserPwd
		return nil, err
	}

	return user, err
}

// Register 处理注册逻辑
func (d *UserDao) Register(user *message.User) error {
	//从连接池取出连接
	conn := d.pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	_, err := d.getUserById(conn, user.UserId)
	if err == nil { //若从数据库取到了数据(返回nil)，则说明要添加的用户已经存在与数据库中
		err = ErrUserExists
		return err
	}

	//这时，说明id在redis还没有，则可以完成注册
	data, err := json.Marshal(user) //序列化
	if err != nil {
		return err
	}

	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("Save registered user error", err)
		return err
	}
	return nil
}
