package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type MessageMan struct {
}

// ReadMessageFromFile 读消息
func (m *MessageMan) ReadMessageFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open ", filePath, "error= ", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("file close error=", err)
		}
	}(file)

	reader := bufio.NewReader(file)
	for {

		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(readString)
	}
	return nil
}

// WriteMessageToFile 写消息
func (m *MessageMan) WriteMessageToFile(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("OpenFile error= ", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("file Close error=", err)
		}
	}(file)

	writer := bufio.NewWriter(file)
	str := content + "\r\n"
	_, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("write content error=", err)
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
