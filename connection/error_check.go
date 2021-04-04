package connection

import "fmt"

// 检查错误
func checkError(err error, info string) (res bool){

	if err != nil {
		fmt.Println(info + " " +  err.Error())
		return false
	}
	return true
}
