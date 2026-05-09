package network

import "net"

func GetLocalIP(conn net.Conn) string {
	addr := conn.LocalAddr().(*net.TCPAddr)
	return addr.IP.String()
}
