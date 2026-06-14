// HAI
package actions

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/anothermeer/sambal/internal/core/device"
	"github.com/anothermeer/sambal/internal/core/network"
	"github.com/anothermeer/sambal/internal/core/protocol"
	"github.com/anothermeer/sambal/internal/core/version"
)

func SendFile(address string, path string) {
	///////////////////////////////////////////   CHECK FILE EXIST AND CONNECT   ///////////////////////////////////////////
	// detect if file exist
	info, err := os.Stat(path)
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

	///////////////////////////////////////////   HELLO   ///////////////////////////////////////////

	// HELLO SERVER
	protocol.Send(conn, protocol.Message{
		Type: "HELLO",
		Payload: map[string]any{
			"version":  version.AppVersion,
			"protocol": version.ProtocolVersion,
			"device": map[string]any{
				"id":   device.GetID(),
				"name": device.GetName(),
			},
		},
	})

	// I HELLO YOU, PLS RESPOND HELLO PLSSSSSSSS
	resp, err := protocol.Receive(conn)
	if err != nil {
		fmt.Println("Failed to receive HELLO_ACK:", err)
		return
	}

	helloPayload, ok := resp.Payload.(map[string]any)
	fmt.Println("Server:", resp.Type)

	// not HELLO
	if resp.Type != "HELLO_ACK" {
		fmt.Println("Expected HELLO_ACK, but got", resp.Type)
		return
	}

	/////////////////////////////////////////// CHECK VERSION   ///////////////////////////////////////////
	if ok {
		fmt.Println("Remote Version:", helloPayload["version"])
		fmt.Println("Remote Protocol:", helloPayload["protocol"])
	}
	remoteProtocol, _ := helloPayload["protocol"].(string)

	if remoteProtocol != strconv.Itoa(version.ProtocolVersion) {
		fmt.Println()
		fmt.Println("WARNING: Protocol mismatch!")
		fmt.Println("Local Protocol :", version.ProtocolVersion)
		fmt.Println("Remote Protocol:", remoteProtocol)
		fmt.Println()
	}

	if remoteProtocol < strconv.Itoa(version.MinProtocolVersion) {
		fmt.Println("ERROR: Remote protocol too old.")
		fmt.Println("Required:", version.MinProtocolVersion)
		fmt.Println("Remote:", remoteProtocol)
		return
	}

	/////////////////////////////////////////// CALC FILE HASH   ///////////////////////////////////////////

	// GONNA CALCULAT DA FILE HASH
	println("Calculating file hash...")
	checksum, err := network.SHA256(path)
	if err != nil {
		fmt.Println("Hash failed:", err)
		return
	}

	fmt.Println("SHA256:", checksum)

	///////////////////////////////////////////   REQUEST SEND FILE   ///////////////////////////////////////////

	// I HAVE A FILE FOR YA, OK OR NO?
	protocol.Send(conn, protocol.Message{
		Type: "FILE_OFFER",
		Payload: protocol.FileOffer{
			Name:   filepath.Base(path),
			Size:   info.Size(),
			SHA256: checksum,
		},
	})

	// FILE ACCEPT?
	resp, err = protocol.Receive(conn)
	if err != nil {
		fmt.Println("Failed to receive FILE_ACCEPT:", err)
		return
	}

	fmt.Println("Server:", resp.Type)

	fileAcceptPayload, ok := resp.Payload.(map[string]any)

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
	if !ok {
		fmt.Println("Invalid FILE_ACCEPT payload")
		return
	}

	// K SO WER UR ADDRES?
	uploadURL, ok := fileAcceptPayload["http_url"].(string)
	if !ok {
		fmt.Println("Missing upload URL")
		return
	}
	fmt.Println("Target Upload URL:", uploadURL)

	///////////////////////////////////////////   SEND FILE   ///////////////////////////////////////////

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
