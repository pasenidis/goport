package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"edpasenidis.tech/goport/internal"
)

func main() {
	var domain string
	var start int
	var end int

	flag.StringVar(&domain, "d", "scanme.nmap.org", "Specify domain. Default is google.com")
	flag.IntVar(&start, "s", 1, "Start of scan range.")
	flag.IntVar(&end, "e", 1024, "End of scan range.")

	flag.Usage = func() {
		fmt.Printf("\nUsage of goport: \n")
		flag.PrintDefaults()
		fmt.Printf("./goport -d example.com\n")
		fmt.Printf("./goport -d example.com -s 1 -e 1024\n")
		fmt.Printf("./goport -d example.com -s 20 -e 777\n")
	}

	flag.Parse()

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
	d := time.Now()
	dS := d.String()

	for index, port := range openports {
		fmt.Printf("%d open\n", port)
		if index == 0 {
			err := ioutil.WriteFile(fmt.Sprintf("goport_%s.txt", dS[:10]), []byte(fmt.Sprintf("website: %s\n%d is open", domain, port)), 0644)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			file, err := os.OpenFile(fmt.Sprintf("goport_%s.txt", dS[:10]), os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			if _, err := file.WriteString(fmt.Sprintf("\n%d is open", port)); err != nil {
				log.Fatal(err)
			}
		}
	}
}
