package main

import (
	"fmt"
	"net"
)

func main() {
	// コネクションを得る
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("error: ", err)
	}

	// ここから読み取り

	data := make([]byte, 1024)
	count, _ := conn.Read(data)
	fmt.Println(string(data[:count]))

	// 出力結果
	// Hello, net pkg!
}
