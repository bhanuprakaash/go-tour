package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func worker(ip string, wg *sync.WaitGroup, ports <-chan int, timeout time.Duration) {
	defer wg.Done()
	for port := range ports {
		address := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			fmt.Printf("%s is OPEN\n", address)
			conn.Close()
		}
	}
}

func main() {
	var timeout time.Duration
	var workers int
	var ip string

	flag.StringVar(&ip, "ip", "127.0.0.1", "ip address")
	flag.DurationVar(&timeout, "timeout", 500*time.Millisecond, "timout(e.g. 500ms, 2s)")
	flag.IntVar(&workers, "workers", 100, "workers")
	flag.Parse()

	if workers <= 0 {
		fmt.Println("workers must be > 0")
		return
	}

	var wg sync.WaitGroup
	var ports = make(chan int, workers)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ip, &wg, ports, timeout)
	}

	for port := 1; port <= 65535; port++ {
		ports <- port
	}
	close(ports)
	wg.Wait()
	fmt.Println("[done]")

}
