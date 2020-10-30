package main

import (
	"fmt"
	"sort"

	"edpasenidis.tech/goport/internal"
)

func main() {
	domain, start, end := internal.Args()

	fmt.Printf("Scanning for %s | %v -> %v\n", domain, start, end)

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := start - 1; i < cap(ports); i++ {
		go internal.Worker(domain, ports, results)
	}

	go func() {
		for i := start; i <= end; i++ {
			ports <- i
		}
	}()

	for i := start - 1; i < end; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)

	internal.Reporter(openports, domain)
}
