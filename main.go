package main

import (
	"fmt"
	"net"

	"github.com/geoffrey-anto/scanner/ping"
	"github.com/geoffrey-anto/scanner/utils"
)

func main() {
	pinger := ping.NewPinger(utils.GetLocalIP(), 3, 5, 100)

	response := pinger.PingAllLocalIP()

	fmt.Printf("%-15s %-15s\n", "Hostname", "IP Address")
	fmt.Println("-----------------------------")
	for _, r := range response {
		if r.Success {
			hostname, _ := net.LookupAddr(r.IP.String())
			if len(hostname) == 0 {
				fmt.Printf("%-15s %-15s\n", "N/A", r.IP.String())
			} else {
				fmt.Printf("%-15s %-15s\n", hostname[0], r.IP.String())
			}
		}
	}
}
