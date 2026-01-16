package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
}

type Result struct {
	id    int
	value int
}

func worker(jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(100 * time.Millisecond)
		results <- Result{id: job.id, value: rand.Intn(100)}
	}
}

func main() {
	const numJobs = 100
	const numWorkers = 3
	var wg sync.WaitGroup

	var jobs = make(chan Job, numJobs)
	var results = make(chan Result, numJobs)

	for range numWorkers {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- Job{id: i}
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result.id, result.value)
	}

	fmt.Println("[DONE]")

}
