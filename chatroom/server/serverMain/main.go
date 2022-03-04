package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"happiness999.cn/chatroom/server/model"
	"happiness999.cn/chatroom/server/process"
	"net"
	"time"
)

//服务器监听9999端口并等待客户端连接,开启协程处理获得的TCP字节流
func main() {
	//初始化redis连接池
	initPool("localhost:6379", 16, 0, time.Second*300)

	//创建UserDao实例
	initUserDao()

	//服务器开启9999端口进行监听
	fmt.Println("My Server Listening ...")
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Server Listen error=", err)
		return
	}
	defer ln.Close()

	//循环等待client连接server
	for {
		fmt.Println("waiting for client connect server ...")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("listener Accept error=", err)
		}

		//连接成功后，启动一个goroutine与client保持通信
		p := &process.Processor{
			Conn: conn,
		}
		go p.Process()
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
	}
}
