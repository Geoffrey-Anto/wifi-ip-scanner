package utils

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func GetLocalIP() net.IP {
	ifaces, err := net.Interfaces()

	if err != nil {
		panic(err)
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip.String()[:3] == "192" {
				return ip
			}
		}
	}

	return nil
}

func GetHostIP(ip net.IP) net.IP {
	return ip.Mask(net.CIDRMask(24, 32))
}

func CheckIP(ip net.IP, waitTime int, count int, timeout int) bool {
	out, _ := exec.Command("ping", ip.String(), fmt.Sprintf("-c %d", waitTime), fmt.Sprintf("-i %d", count), fmt.Sprintf("-w %d", timeout)).Output()
	if strings.Contains(string(out), "Destination Host Unreachable") {
		return false
	} else {
		return true
	}
}
