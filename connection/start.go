package connection

import (
	"fmt"
)

func StartTCP(params []string) {
	length := len(params)

	if length != 2 {
		fmt.Println("wrong params")
	} else if params[0] == "server" && length == 2 {
		startServer(params[1])
	} else if params[0] == "client" && length == 2 {
		startClient(params[1])
	} else {
		fmt.Println("wrong params")
	}
}