package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Simple TCP port scanner")

	var host string
	flag.StringVar(&host, "host", "localhost", "hostname of a scanned host")

	var portRange string
	flag.StringVar(&portRange, "portrange", "1-1000", "ranged of ports to scan")

	var ports string
	flag.StringVar(&ports, "ports", "21,22,23,80,443", "specific ports to scan")

	var specific bool
	flag.BoolVar(&specific, "specific", false, "scan specific ports instead of range")

	flag.Parse()

	ipAddrs, _ := net.LookupIP(host)

	fmt.Println("Host:", host)
	fmt.Println("Addresses:", ipAddrs)
	fmt.Println("Port range:", portRange)
	fmt.Println("Ports:", ports)
	fmt.Println("Specific:", specific)
	fmt.Println()

	for a := 0; a < len(ipAddrs); a++ {
		if specific {
			portsToScan := strings.Split(ports, ",")

			for i := 0; i < len(portsToScan); i++ {
				checkPortStatus(ipAddrs[a].String(), portsToScan[i])
			}

		} else {
			splitRange := strings.Split(portRange, "-")
			lowPort, err := strconv.ParseUint(splitRange[0], 10, 16)

			if err != nil {
				log.Fatal(err)
			}

			highPort, err := strconv.ParseUint(splitRange[1], 10, 16)

			if err != nil {
				log.Fatal(err)
			}

			for i := lowPort; i <= highPort; i++ {
				strPort := fmt.Sprint(i)
				go checkPortStatus(ipAddrs[a].String(), strPort)
			}
		}
	}
}

func checkPortStatus(host, port string) bool {
	addr := net.JoinHostPort(host, port)

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return false
	} else {
		conn.Close()
		fmt.Printf("%v: OPEN (%v)\n", port, host)
		return true
	}
}
