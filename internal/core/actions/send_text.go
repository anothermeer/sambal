package actions

import (
	"fmt"
	"net"

	"github.com/anothermeer/sambal/internal/core/protocol"
)

func SendText(address string, text string) {
	conn, _ := net.Dial("tcp", address)
	defer conn.Close()

	protocol.Send(conn, protocol.Message{Type: "HELLO"})
	protocol.Receive(conn)

	protocol.Send(conn, protocol.Message{
		Type: "TEXT_SEND",
		Payload: protocol.TextTrans{
			Text: text,
		},
	})

	fmt.Println("Text sent")
}
