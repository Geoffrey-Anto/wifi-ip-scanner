package ping

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/geoffrey-anto/wifi-ip-scanner/utils"
)

type PingResponse struct {
	IP       net.IP
	Success  bool
	Hostname string
}

type Pinger struct {
	IP net.IP

	waitTime int
	count    int
	timeout  int
}

func NewPinger(ip net.IP, waitTime int, count int, timeout int) *Pinger {
	return &Pinger{
		IP:       ip,
		waitTime: waitTime,
		count:    count,
		timeout:  timeout,
	}
}

func (p *Pinger) PingAllLocalIP() []PingResponse {
	hostIp := utils.GetHostIP(p.IP)

	ipsFound := 0

	wg := sync.WaitGroup{}

	t := time.Now()

	responses := make([]PingResponse, 0)

	for i := 1; i < 255; i++ {
		// CIDR Mask 24
		// TODO: Make this more dynamic
		host := strings.Join(strings.Split(hostIp.String(), ".")[0:3], ".")

		addr := host + "." + fmt.Sprintf("%d", i)

		wg.Add(1)

		go func() {
			defer wg.Done()

			if utils.CheckIP(net.ParseIP(addr), p.waitTime, p.count, p.timeout) {
				ipsFound++
				hostnames, _ := net.LookupAddr(addr)
				hostname := "N/A"

				if len(hostnames) > 0 {
					hostname = hostnames[0]
				}

				responses = append(responses, PingResponse{
					IP:       net.ParseIP(addr),
					Success:  true,
					Hostname: hostname,
				})
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Found %d ips\n", ipsFound)
	fmt.Printf("Local IP: %s\n", p.IP.String())
	fmt.Printf("Time taken: %s\n", time.Since(t))

	return responses
}
