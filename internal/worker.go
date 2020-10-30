package internal

import (
	"fmt"
	"net"
)

// Worker scans for ports
func Worker(domain string, ports chan int, results chan int) {
	// gets called by the main package
	for p := range ports {
		address := fmt.Sprintf("%s:%d", domain, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
