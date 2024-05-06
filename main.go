package main

import (
	"fmt"

	"github.com/geoffrey-anto/wifi-ip-scanner/ping"
	"github.com/geoffrey-anto/wifi-ip-scanner/utils"
)

func main() {
	pinger := ping.NewPinger(utils.GetLocalIP(), 3, 5, 100)

	response := pinger.PingAllLocalIP()

	fmt.Printf("%-35s %-35s\n", "Hostname", "IP Address")
	fmt.Println("--------------------------------------------------------")
	for _, r := range response {
		fmt.Printf("%-35s %-35s\n", r.Hostname, r.IP.String())
	}
}
