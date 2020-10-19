package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//Get the args without the executable
	argsWithoutProg := os.Args[1:]
	//get the number of args excluding the trailing output file
	end := len(argsWithoutProg) - 1

	//get the slices of the args for the inputs and outputs
	inputs := argsWithoutProg[:end]
	outfile := argsWithoutProg[end]

	// For write access and to create the file if it doesn't exist
	of, err := os.Create(outfile) 
	if err != nil {
		fmt.Println(err)
		return
	}
	defer of.Close()
	for _, s := range inputs {
		infile, err := os.Open(s) // For read access.
		if err != nil {
			log.Fatal(err)
			return
		}
		defer if.Close()
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
