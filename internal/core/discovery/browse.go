package discovery

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/libp2p/zeroconf/v2"
)

func Browse(timeout time.Duration) ([]Device, error) {
	entries := make(chan *zeroconf.ServiceEntry)
	var devices []Device

	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	// now here's a goroutine
	go func() {
		defer wg.Done()
		for entry := range entries {
			d := Device{
				Name:     entry.Instance,
				Port:     entry.Port,
				LastSeen: time.Now(),
			}

			// pick first IPv4 if available
			if len(entry.AddrIPv4) > 0 {
				d.Addr = entry.AddrIPv4[0]
			} else if len(entry.AddrIPv6) > 0 {
				d.Addr = entry.AddrIPv6[0]
			}

			// parse TXT records
			for _, txt := range entry.Text {
				parts := strings.SplitN(txt, "=", 2)
				if len(parts) != 2 {
					continue
				}

				switch parts[0] {
				case "id":
					d.ID = parts[1]

				case "version":
					d.Version = parts[1]

				case "protocol":
					p, _ := strconv.Atoi(parts[1])
					d.Protocol = p

				case "name":
					d.Name = parts[1]

				case "port":
					p, _ := strconv.Atoi(parts[1])
					d.Port = p
				}

			}

			found := false
			for i := range devices {
				if devices[i].ID == d.ID {
					devices[i] = d
					found = true
					break
				}
			}

			if !found {
				devices = append(devices, d)
			}
		}
	}()

	err := zeroconf.Browse(ctx, Service, Domain, entries, nil)
	if err != nil {
		return nil, err
	}

	<-ctx.Done()
	wg.Wait()

	return devices, nil
}
