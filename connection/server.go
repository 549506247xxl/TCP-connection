package connection

import (
	"fmt"
	"net"
)

var (
	recBuffer = 350000000
)

// 服务端接受数据线程
func  handler(conn net.Conn) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	buf := make([]byte, recBuffer)
	for {
		// read data from this connection
		length, err := conn.Read(buf)
		if checkError(err, "Connection") == false {
			_ = conn.Close()
			break
		}
		if length > 0 {
			buf[length] = 0
		}
		// 打印从TCP连接客户端发送过来的数据
		//fmt.Println(string(buf[0:length]))
	}

}


// 启动服务器

func startServer(port string){
	// 开启服务的IP地址未定义，监听所有单播和任播的IP地址
	service := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "ResolveTCPAddr")

	l, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err, "ListenTCP")

	fmt.Println("Start TCP Server Success")


	for {
		conn, err := l.Accept()
		checkError(err, "Accept")
		fmt.Println("Accepting ", conn.RemoteAddr().String())
		go handler(conn)
	}
}