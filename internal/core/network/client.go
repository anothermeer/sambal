package network

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"

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
	fmt.Println("Server:", resp.Type)

	protocol.Send(conn, protocol.Message{
		Type: "FILE_OFFER",
		Payload: protocol.FileOffer{
			Name: "test.txt",
			Size: 2113,
		},
	})

	resp, _ = protocol.Receive(conn)
	fmt.Println("Server:", resp.Type)

	protocol.Send(conn, protocol.Message{
		Type: "FILE_START",
		Payload: map[string]any{
			"url": "http://localhost:3722/upload",
		},
	})

	uploadFile("test.txt")
}

func uploadFile(path string) {
	file, _ := os.ReadFile(path)

	req, _ := http.NewRequest("POST", "http://localhost:3722/upload", bytes.NewReader(file))
	req.Header.Set("X-Sambal-Name", path)

	http.DefaultClient.Do(req)
}
