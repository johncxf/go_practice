// 建立 TCP 连接来实现简单的 HTTP 协议
package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

// 检查网络连接
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// 读取响应数据
func readFully(conn net.Conn) ([]byte, error) {
	// 读取所有响应数据后主动关闭连接
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	// 从参数中读取主机信息
	service := os.Args[1]

	// 建立网络连接
	conn, err := net.Dial("tcp", service)
	// 连接出错则打印错误消息并退出程序
	checkError(err)

	// 调用返回的连接对象提供的 Write 方法发送请求
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// 通过连接对象提供的 Read 方法读取所有响应数据
	result, err := readFully(conn)
	checkError(err)

	// 打印响应数据
	fmt.Println(string(result))

	os.Exit(0)
}
