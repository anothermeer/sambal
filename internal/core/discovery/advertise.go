package discovery

import (
	"fmt"
	"strconv"

	"github.com/grandcat/zeroconf"

	"github.com/anothermeer/sambal/internal/core/device"
	"github.com/anothermeer/sambal/internal/core/version"
)

var server *zeroconf.Server

func StartAdvertiser() error {
	var err error

	server, err = zeroconf.Register(
		device.GetName(),
		"_sambal._tcp",
		"local.",
		3721,
		[]string{
			"id=" + device.GetID(),
			"protocol=" + strconv.Itoa(version.ProtocolVersion),
			"version=" + version.AppVersion,
		},
		nil,
	)

	zeroconf.Register("SambalTest", "_http._tcp", "local.", 3721, nil, nil)

	fmt.Printf(
		"[DBG] mDNS advertising: %s._sambal._tcp.local\n",
		device.GetName(),
	)
	if err != nil {
		fmt.Println("mDNS Register Error:", err)
		return err
	}

	fmt.Println("mDNS advertising:", device.GetName())

	return err
}

func StopAdvertiser() {
	if server != nil {
		server.Shutdown()
	}
}
