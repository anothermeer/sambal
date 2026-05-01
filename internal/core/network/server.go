package network

import (
	"encoding/json"
	"fmt"
	"net"

	"context"
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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	msg, err := protocol.Receive(conn)
	if err != nil {
		fmt.Println("Error:", err)
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
	}

}
