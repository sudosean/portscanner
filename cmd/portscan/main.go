package main

import (
	"fmt"
	"net"
	"sort"
)

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	/*
		send to the workers in a separate goroutine
		because the results-gathering loop (below) neds to start
		before more than 100 items of work can continue
	 */
	go func() {
		for i := 0; i <= 65535; i++ {
			ports <- i
		}
	}()


	/*
		receives the results and checks if it's a port
		if so, we add to the slice
	 */
	for i := 0; i < 65535; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}
// worker accepts two channels. if the port is
// closed, send a zero (0) and if it's open send the port
func worker(ports, results chan int){
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		/*
			separate channel to communicate results
			from worker to main thread
		 */
		results <- p
	}
}
