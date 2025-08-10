package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
	"time"
)

//Here is the list of ports to be scanned.
//If you want to scan an additional port, please add it to the list below.
var commonPorts = map[int]string{
	21:   "FTP",
	22:   "SSH",
	23:   "Telnet",
	30:	  "Telnet-Alt",
	25:   "SMTP",
	53:   "DNS",
	80:   "HTTP",
	110:  "POP3",
	143:  "IMAP",
	443:  "HTTPS",
	465:  "SMTPS",
	993:  "IMAPS",
	995:  "POP3S",
	3306: "MySQL",
	3389: "RDP",
	5432: "PostgreSQL",
	5900: "VNC",
	6379: "Redis",
	8080: "HTTP-Alt",
	8443: "HTTPS-Alt",
}

func scanPort(protocol, hostname string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	host := flag.String("host", "", "Target hostname or IP address")
	timeout := flag.Int("timeout", 500, "Timeout in milliseconds")
	flag.Parse()

	if *host == "" {
		fmt.Println("Please provide a host using -host flag.")
		os.Exit(1)
	}

	fmt.Println("Scanning", *host)
	fmt.Println("-------------------------")

	timeoutDuration := time.Duration(*timeout) * time.Millisecond
	var ports []int
	for p := range commonPorts {
		ports = append(ports, p)
	}
	sort.Ints(ports)

	for _, port := range ports {
		open := scanPort("tcp", *host, port, timeoutDuration)
		if open {
			fmt.Printf("[+] Port %d (%s): OPEN\n", port, commonPorts[port])
		} else {
			fmt.Printf("[-] Port %d (%s): closed\n", port, commonPorts[port])
		}
	}
}
