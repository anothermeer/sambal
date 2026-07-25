package discovery

import (
	"net"
	"time"
)

type Device struct {
	ID       string
	Name     string
	Addr     net.IP
	Port     int
	Version  string
	Protocol int
	LastSeen time.Time
}
