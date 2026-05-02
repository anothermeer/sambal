package actions

import (
	"fmt"
	"net"

	"github.com/anothermeer/sambal/internal/core/network"
	"github.com/anothermeer/sambal/internal/core/protocol"
	"github.com/anothermeer/sambal/internal/core/version"
)

func SendFile(address string, path string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connection Failed:", err)
		return
	}
	defer conn.Close()

	protocol.Send(conn, protocol.Message{
		Type: "HELLO",
		Payload: map[string]any{
			"version": version.Version,
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
			Name: path,
			Size: 0, // TODO: later fix this ah
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

	network.UploadFile("test.txt")
}
