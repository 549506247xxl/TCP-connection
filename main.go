package main

import (
	"fmt"
	"github.com/549506247xxl/TCP_golang/connection"
	"github.com/549506247xxl/TCP_golang/monitor"
	"os"
)

// 声明全局变量
var (
	serverIP = "127.0.0.1"  // 服务端IP地址
	//serverIP = "47.94.104.34"  // 服务端IP地址
	port     = "8087"       // 服务端端口号
	device   = "ens33"         // 监听的网卡
	//device   = "eth0"         // 监听的网卡
)

func main() {
	way := os.Args[1]
	if way == "connection" || way == "-c" || way == "conn" {
		var params []string
		if os.Args[2] == "server" {
			params = append(params, "server", port)
		} else if os.Args[2] == "client" {
			params = append(params, "client", serverIP+":"+port)
		} else {
			fmt.Println("params error")
			fmt.Println("expected client server")
		}
		connection.StartTCP(params)
	} else if way == "monitor" || way == "-m" {
		monitor.Capture(device)
	} else {
		fmt.Println("params error")
	}
}
