package discovery

import (
	"fmt"
	"strconv"

	"github.com/libp2p/zeroconf/v2"

	"github.com/anothermeer/sambal/internal/core/device"
	"github.com/anothermeer/sambal/internal/core/version"
)

var server *zeroconf.Server

func StartAdvertiser() error {
	var err error

	txt := []string{
		"id=" + device.GetID(),
		"name=" + device.GetName(),
		"version=" + version.AppVersion,
		"protocol=" + strconv.Itoa(version.ProtocolVersion),
		"port=" + strconv.Itoa(DefaultPort),
	}

	server, err = zeroconf.Register(
		device.GetName(),
		Service,
		Domain,
		DefaultPort,
		txt,
		nil,
	)
	if err != nil {
		return fmt.Errorf("register mdns service: %w", err)
	}

	fmt.Printf(
		"[DBG] mDNS advertising: %s.%s.%s\n",
		device.GetName(),
		Service,
		Domain,
	)

	return nil
}

func StopAdvertiser() {
	if server != nil {
		server.Shutdown()
		server = nil
	}
}
