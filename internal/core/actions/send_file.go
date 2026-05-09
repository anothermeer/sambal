// HAI
package actions

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/anothermeer/sambal/internal/core/network"
	"github.com/anothermeer/sambal/internal/core/protocol"
	"github.com/anothermeer/sambal/internal/core/version"
)

func SendFile(address string, path string) {
	// detect if file exist
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("File not found:", path)
		return
	}

	// try to establish connection
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Connection Failed:", err)
		return
	}
	defer conn.Close()

	info, _ := os.Stat(path)

	// HELLO
	protocol.Send(conn, protocol.Message{
		Type: "HELLO",
		Payload: map[string]any{
			"version": version.Version,
			"device": map[string]any{
				"id":   "test1",
				"name": "test-device1",
			},
		},
	}) // not HELLO
	resp, err := protocol.Receive(conn)
	if err != nil {
		fmt.Println("Failed to receive HELLO_ACK:", err)
		return
	}
	fmt.Println("Server:", resp.Type)

	//fmt.Printf("RAW RESPONSE: %#v\n", resp)
	//fmt.Printf("PAYLOAD TYPE: %T\n", resp.Payload)

	// I KNOW YOU HELLO, SO HELLO YOU
	if resp.Type != "HELLO_ACK" {
		fmt.Println("Expected HELLO_ACK, but got", resp.Type)
		return
	}

	// I HAVE A FILE FOR YA, OK OR NO?
	protocol.Send(conn, protocol.Message{
		Type: "FILE_OFFER",
		Payload: protocol.FileOffer{
			Name: filepath.Base(path),
			Size: info.Size(),
		},
	})

	resp, err = protocol.Receive(conn)
	fmt.Println("Server:", resp.Type)

	// WHERE R U??
	if err != nil {
		fmt.Println("Failed to receive FILE_ACCEPT:", err)
		return
	}
	// WHA- YOU SAY NO??? 💔
	if resp.Type == "FILE_DECLINE" {
		fmt.Println("Target device declined file transfer.")
		return
	}
	// WHAT DO YOU WANT??? 😭
	if resp.Type != "FILE_ACCEPT" {
		fmt.Println("Expected FILE_ACCEPT, got", resp.Type)
		return
	}

	// WTF YOU SAY???
	payload, ok := resp.Payload.(map[string]any)
	if !ok {
		fmt.Println("Invalid FILE_ACCEPT payload")
		return
	}

	// K SO WER UR ADDRES?
	uploadURL, ok := payload["http_url"].(string)
	if !ok {
		fmt.Println("Missing upload URL")
		return
	}
	fmt.Println("Target Upload URL:", uploadURL)

	// R U READY?
	protocol.Send(conn, protocol.Message{
		Type: "FILE_START",
	})

	// MIKU MIKU BEAAAAAAAAAAAM!!!
	err = network.UploadFile(path, uploadURL)
	if err != nil { // OH NOEZ UPLOAD FAILD
		fmt.Println("Upload failed:", err)
		return
	}
}

// KTHXBAI
