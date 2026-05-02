package actions

import (
	"fmt"
	"net"

	"github.com/anothermeer/sambal/internal/core/protocol"
)

func Ping(address string) {
	conn, _ := net.Dial("tcp", address)
	defer conn.Close()

	protocol.Send(conn, protocol.Message{Type: "HELLO"})
	resp, _ := protocol.Receive(conn)

	fmt.Println("Response:", resp.Type)
}
