package network

import (
	"fmt"
	"net"

	"github.com/anothermeer/sambal/internal/core/protocol"
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

func SendHello(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connection Failed:", err)
		return
	}
	defer conn.Close()

	protocol.Send(conn, protocol.Message{
		Type: "HELLO",
		Payload: map[string]any{
			"version": "v0.1",
			"device": map[string]any{
				"id":   "test1",
				"name": "test-device1",
			},
		},
	})

	resp, _ := protocol.Receive(conn)
	fmt.Println("Server responded:", resp.Type)
}
