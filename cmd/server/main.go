package main

import (
	"fmt"
	"net"
)

func main() {
	// コネクションを得る
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("cannot listen", err)
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("cannot accept", err)
	}

	// ここから送信

	str := "Hello, net pkg!"
	data := []byte(str)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("cannot write", err)
	}
}
