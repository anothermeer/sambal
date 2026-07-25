package discovery

import (
	"context"
	"fmt"
	"time"

	"github.com/grandcat/zeroconf"
)

func ListDevices() {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		fmt.Println("mDNS resolver error:", err)
		return
	}

	entries := make(chan *zeroconf.ServiceEntry)

	go func() {
		for entry := range entries {
			fmt.Println("[Discovery] Found:", entry.Instance)
			fmt.Println("[DBG] Instance:", entry.Instance)
			fmt.Println("[DBG] Host:", entry.HostName)
			fmt.Println("[DBG] IPv4:", entry.AddrIPv4)
			fmt.Println("[DBG] Port:", entry.Port)
			fmt.Println("[DBG] TXT:", entry.Text)
		}
	}()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	fmt.Println("[DBG] Starting browse...")
	fmt.Println("Browsing for _sambal._tcp.local")
	err = resolver.Browse(
		ctx,
		"_sambal._tcp",
		"local.",
		entries,
	)
	fmt.Println("[DBG] Browse started")

	if err != nil {
		fmt.Println("mDNS Browse error:", err)
		return
	}

	fmt.Println("[DBG] Waiting for discoveries...")
	<-ctx.Done()
	fmt.Println("[DBG] Browse timeout reached")
}
