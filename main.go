package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)


func scanPort(host string, port int, timeout time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return
	}
	conn.Close()
	fmt.Printf("Port %d: Is Open\n", port)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Welcome to GoNeMap by eauzun\nUsage: go run main.go <target>")
		return
	}

	host := os.Args[1]
	timeout := 500 * time.Millisecond

	var wg sync.WaitGroup

	fmt.Println("Scanning...")

	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go scanPort(host, port, timeout, &wg)
	}

	wg.Wait()

	fmt.Println("Scan Process Completed")
}
