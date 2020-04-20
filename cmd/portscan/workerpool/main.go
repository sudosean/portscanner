package portscanner

import (
	"fmt"
	"sync"
)

func WorkerPool() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports) ; i++ {
		go worker(ports, &wg)
	}
	for i := 0; i < 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

func worker(ports chan int, wg *sync.WaitGroup){
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}
