package connection

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sync"
)

// 输入$send指令，可以发送下面整个文件的内容
// 文件路径是一个重要问题
// 需要打开的文件要和main.go放在同一个目录下
var filename = "xx.txt"

// 客户端发送线程

func sendMessage(conn net.Conn, wg *sync.WaitGroup){
	var input string
	username := conn.LocalAddr().String()
	for {
		// use this way to read string when meet '\n' to stop rather than ' '
		// but this will append '\n' at the end of [input]
		inputReader := bufio.NewReader(os.Stdin)
		input, _ = inputReader.ReadString('\n')
		input = input[0: len(input)-1]

		if input == "quit"{
			fmt.Println("ByeBye ...")
			_ = conn.Close()
			wg.Done()
			os.Exit(0)
		} else if input == "$send" {
			fmt.Println("a file will be sent.")
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Println("File reading error!")
				return
			}
			// 将从文件中读取的内容，转换为string类型
			input = string(data)
		}

		_, err := conn.Write([]byte(username + " Say ::: " + input))
		if err != nil{
			fmt.Println(err.Error())
			_ = conn.Close()
			wg.Done()
			break
		}
	}
}

// 客户端启动函数
// addr = ip:port  想要连接的服务器的地址和端口
func startClient(addr string){
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err, "ResolveTCPAddr")

	// a local address is automatically chosen if the second param laddr is nil
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(err, "DialTCP")

	var wg sync.WaitGroup
	wg.Add(1)
	go sendMessage(conn, &wg)
	wg.Wait()
}