package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Reporter saves logs to text files
func Reporter(openports []int, domain string) {
	d := time.Now()
	dS := d.String()

	for index, port := range openports {
		fmt.Printf("%d open\n", port)
		filename := fmt.Sprintf("goport_%s.txt", dS[:10])

		if index == 0 {
			err := ioutil.WriteFile(filename, []byte(fmt.Sprintf("website: %s\n%d is open", domain, port)), 0644)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
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
