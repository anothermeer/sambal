package protocol

import (
	"encoding/json"
	"net"
)

func Send(conn net.Conn, msg Message) error {
	encoder := json.NewEncoder(conn)
	return encoder.Encode(msg)
}
