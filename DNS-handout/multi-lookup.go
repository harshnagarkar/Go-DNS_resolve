package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

func request(name string, ch chan string) {

	fmt.Println("test2")
	infile, err := os.Open(name) // For read access.
	if err != nil {
		log.Fatal(err)
		return
	}
	defer infile.Close()
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		var hostname = scanner.Text()
		// resolve(filename)
		fmt.Println(hostname)
		ch <- hostname
	}
}

func resolve(filename string, ch chan string) {
	// // fmt.Println("Results: "+filename)

	// defer of.Close()

	for {
		// if(len(ch)>0){
		// fmt.Println("hello")

		var hostname = <-ch
		// hostname<-ch
		ips, err := net.LookupIP(hostname)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dnslookup error: %s\n", hostname)
		}

		of, err := os.Open(filename) // For read access.
		if err != nil {
			log.Fatal(err)
			return
		}

		if len(ips) > 0 {
			fmt.Println(hostname + " resolve: " + ips[0].String())
			//pass to the channel
			fmt.Fprintf(of, "%s %s\n", hostname, ips[0].String())
		} else {
			fmt.Fprintf(of, "%s \n", hostname)
			//pass to the channel
		}
		of.Close()
		// }
	}
}

func main() {
	ch := make(chan string)
	//Get the args without the executable
	argsWithoutProg := os.Args[1:]
	//get the number of args excluding the trailing output file
	end := len(argsWithoutProg) - 1

	//get the slices of the args for the inputs and outputs
	// inputs := argsWithoutProg[:end]
	outfile := argsWithoutProg[end]

	// For write access and to create the file if it doesn't exist
	of, err := os.Create(outfile)

	if err != nil {
		fmt.Println(err)
		return
	}
	of.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("test")
	go request("input/names1.txt", ch)
	go resolve(outfile, ch)
	wg.Wait()
	// for _, s := range inputs {
	// 	ch <- s
	// }

}
