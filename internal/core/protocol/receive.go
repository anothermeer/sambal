package protocol

import (
	"encoding/json"
	"net"
)

func Receive(conn net.Conn) (Message, error) {
	var msg Message
	decoder := json.NewDecoder(conn)
	err := decoder.Decode(&msg)
	return msg, err
}
