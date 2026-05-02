package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/anothermeer/sambal/internal/core/protocol"
)

func StartTCPSrv(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	// http server for transfering files
	go startHTTPServer()

	fmt.Println("Sambal listening on port", port)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		fmt.Println("\nShutting down server...")
		ln.Close() // this will unblock Accept()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("ERR: ", err)
				continue
			}
		}

		go handleConnection(conn)
	}
}

func startHTTPServer() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		name := r.Header.Get("X-Sambal-Name")

		file, _ := os.Create(name)
		defer file.Close()

		io.Copy(file, r.Body)

		fmt.Println("Received file:", name)
	})

	fmt.Println("HTTP Server listening on port 3722")
	http.ListenAndServe(":3722", nil)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		msg, err := protocol.Receive(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client Disconnected")
			} else {
				fmt.Println("Error:", err)
			}
			return
		}

		switch msg.Type {

		case "HELLO":
			protocol.Send(conn, protocol.Message{
				Type: "HELLO_ACK",
			})

		case "FILE_OFFER":
			fmt.Println("Incoming file offer")

			data, _ := json.Marshal(msg.Payload)

			var offer protocol.FileOffer
			json.Unmarshal(data, &offer)

			fmt.Println("File:", offer.Name, "| Size:", offer.Size)

			// auto accept file offer for now
			protocol.Send(conn, protocol.Message{
				Type: "FILE_ACCEPT",
			})

		case "FILE_START":
			fmt.Println("Client starting upload...")

		case "TEXT_SEND":
			data, _ := json.Marshal(msg.Payload)

			var text protocol.TextTrans
			json.Unmarshal(data, &text)

			fmt.Println("Text received:", text.Text)

		case "CLIPBOARD_SYNC":
			// https://github.com/tiagomelo/go-clipboard
			data, _ := json.Marshal(msg.Payload)

			var clip protocol.CBTrans
			json.Unmarshal(data, &clip)

			fmt.Println("Clipboard received:", clip.Text)
		}
	}
}
