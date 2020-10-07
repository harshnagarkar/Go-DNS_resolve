package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	argsWithoutProg := os.Args[1:]

	end := len(argsWithoutProg) - 1
	inputs := argsWithoutProg[:end]
	outfile := argsWithoutProg[end]
	of, err := os.Create(outfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range inputs {
		infile, err := os.Open(s) // For read access.
		if err != nil {
			log.Fatal(err)
			return
		}
		scanner := bufio.NewScanner(infile)
		for scanner.Scan() {
			ips, err := net.LookupIP(scanner.Text())
			if err != nil {
				fmt.Fprintf(os.Stderr, "dnslookup error: %s\n", scanner.Text())
			}
			if len(ips) > 0 {
				fmt.Fprintf(of, "%s %s\n", scanner.Text(), ips[0].String())
			} else {
				fmt.Fprintf(of, "%s \n", scanner.Text())
			}
		}
	}
}
