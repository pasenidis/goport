package internal

import (
	"flag"
	"fmt"
)

// Args parses CLI arguments and returns them
func Args() (string, int, int) {
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

	return domain, start, end
}
