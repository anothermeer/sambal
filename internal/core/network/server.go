package network

import (
	"fmt"
	"net"

	"context"
	"os/signal"
	"syscall"
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

	fmt.Println("New Connection from ", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	fmt.Println("Received:", string(buffer[:n]))
}
