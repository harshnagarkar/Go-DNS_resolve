package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"sync"
)

type safe_channel struct {
	mu      sync.Mutex
	of_file *os.File
	ch      chan string
	wg      *sync.WaitGroup
}

func request(name string, ch chan string, wg *sync.WaitGroup) {

	infile, err := os.Open(name) // For read access.
	if err != nil {
		log.Fatal(err)
		return
	}
	defer infile.Close()
	defer wg.Done()
	scanner := bufio.NewScanner(infile)

	for scanner.Scan() {
		var hostname = scanner.Text()
		// resolve(filename)
		fmt.Println(hostname)

		ch <- hostname
	}
}

func resolve(c *safe_channel) {
	// // fmt.Println("Results: "+filename)

	// defer of.Close()

	defer c.wg.Done()
	for i := range c.ch {
		// if(len(ch)>0){
		// fmt.Println("hello")

		var hostname = i
		// hostname<-ch
		ips, err := net.LookupIP(hostname)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dnslookup error: %s\n", hostname)
		}

		// of, err := os.(filename) // For read access.
		// if err != nil {
		// 	log.Fatal(err)
		// 	return
		// }

		//Go example 9
		// struct element file pointer and mutex
		// net.
		c.mu.Lock()
		if len(ips) > 0 {
			fmt.Println(hostname + " resolve: " + ips[0].String())
			//pass to the channel
			data := hostname
			// data+=hostname
			for i := 0; i < len(ips); i++ {
				data = data + "," + ips[i].String()
			}
			data += "\n"
			// fmt.Fprintf(c.of_file, "%s %s\n", hostname, ips[0].String())
			fmt.Fprint(c.of_file, data)
		} else {
			fmt.Fprintf(c.of_file, "%s \n", hostname)
			//pass to the channel
		}
		c.mu.Unlock()

	}
}

func main() {
	ch := make(chan string, 5)
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
	// of.Close()

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	safe_ch := safe_channel{of_file: of, ch: ch, wg: &wg2}

	fmt.Println("test")

	wg1.Add(len(inputs))
	for _, s := range inputs {
		// ch <- s
		fmt.Println(s)
		go request(s, ch, &wg1)
	}
	fmt.Printf("The number of cpu are %d", runtime.NumCPU())

	wg2.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go resolve(&safe_ch)
	}

	wg1.Wait()

	close(ch)

}
