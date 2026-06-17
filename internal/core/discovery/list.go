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
		}
	}()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	err = resolver.Browse(
		ctx,
		"_sambal._tcp",
		"local.",
		entries,
	)

	if err != nil {
		fmt.Println("mDNS Browse error:", err)
		return
	}

	<-ctx.Done()
}
