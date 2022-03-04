package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"happiness999.cn/chatroom/client/utils/message"
	"io"
	"net"
)

// Transfer 关联方法到Transfer结构体:定义了conn和缓冲
type Transfer struct {
	Conn net.Conn
	Buf  [1024 * 4]byte //缓冲:用于接收传输过来的字节流
}

// WritePkg 处理数据包并write
func (tf *Transfer) WritePkg(data []byte) error {
	//获取data的length发送给server
	//1,先获取到data的长度 -> 转化成一个表示长度的byte切片
	var pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)

	//发送data的长度
	n, err := tf.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write error=", err)
		return err
	}

	//fmt.Println("server send message length success, length = ", len(data), "data = ", string(data))

	//发送data本身
	n, err = tf.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return err
	}

	return err
}

// ReadPkg 处理数据包并返回Message，Err
func (tf *Transfer) ReadPkg() (message.Message, error) {

	var mes message.Message
	//获取客户端传过来的message length
	//buf := make([]byte, 1024*4)
	//fmt.Println("read data")
	n, err := tf.Conn.Read(tf.Buf[:4])
	if err == io.EOF || err != nil {
		fmt.Println("Client exit session")
		return mes, err
	}

	//将message length 转换为 uint32类型并根据length读取message 内容
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(tf.Buf[:4])
	n, err = tf.Conn.Read(tf.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read package body error")
		return mes, err
	}

	err = json.Unmarshal(tf.Buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("json Unmarshal error")
		return mes, err
	}

	return mes, nil
}
