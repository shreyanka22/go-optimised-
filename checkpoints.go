package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, checkpoint, resume chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d waiting for checkpoint\n", id)
	checkpoint <- true
	<-resume
	fmt.Printf("Worker %d resuming\n", id)
}

func main() {
	numWorkers := 5
	checkpoint, resume := make(chan bool), make(chan bool)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, checkpoint, resume, &wg)
		<-checkpoint
	}

	fmt.Println("All workers reached the checkpoint\nResuming all workers")

	for i := 0; i < numWorkers; i++ {
		resume <- true
	}

	wg.Wait()
	fmt.Println("All workers completed their work")
}
