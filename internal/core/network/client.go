package network

import (
	"fmt"
	"net"
)

func SendMsg(address string, message string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connection Failed:", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte(message))
	fmt.Println("Sent:", message)
}
