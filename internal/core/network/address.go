package network

import "net"

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "172.0.0.1"
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		if ipnet.IP.IsLoopback() {
			continue
		}

		ipv4 := ipnet.IP.To4()
		if ipv4 == nil {
			continue // skip IPv6 for now
		}

		return ipv4.String()
	}

	return "127.0.0.1"
}
