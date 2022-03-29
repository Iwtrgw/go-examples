package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := Decode(reader)
		// 如果客户端关闭，则退出本协程
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("reader.Read error :", err)
			break
		}
		fmt.Printf("received data：%s\n\n", msg)
	}
}

// Decode 拆包
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(2) // 读取包头的前两个字节
	lengthBuf := bytes.NewBuffer(lengthByte)
	var length int16
	// 读取包的实际长度
	err := binary.Read(lengthBuf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int16(reader.Buffered()) < length+2 {
		return "", err
	}
	realData := make([]byte, int(length+2))
	_, err = reader.Read(realData)
	if err != nil {
		return "", err
	}
	return string(realData[2:]), nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("net.Listen.error:", err)
		return
	}

	defer listen.Close()
	fmt.Println("Server Start.....")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Listen.Accept.error:", err)
			continue
		}
		go process(conn)
	}
}
