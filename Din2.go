package main

import (
    "fmt"
    "sync"
    "time"
)

type Chopstick struct {
    sync.Mutex
}

func main() {
    var numPhilosophers, maxDiningCycles int

    fmt.Print("Enter the number of philosophers: ")
    fmt.Scan(&numPhilosophers)

    fmt.Print("Enter the maximum dining cycle count: ")
    fmt.Scan(&maxDiningCycles)

    var wg sync.WaitGroup
    chopsticks := make([]Chopstick, numPhilosophers)
    diningCycles := make([]int, numPhilosophers)

    for i := 0; i < numPhilosophers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Philosopher %d is thinking\n", id)

            for diningCycles[id] < maxDiningCycles {
                time.Sleep(time.Second)
                left := id
                right := (id + 1) % numPhilosophers

                chopsticks[left].Lock()
                chopsticks[right].Lock()

                fmt.Printf("Philosopher %d is eating (Cycle %d)\n", id, diningCycles[id])
                time.Sleep(time.Second)

                chopsticks[left].Unlock()
                chopsticks[right].Unlock()

                diningCycles[id]++
            }

            fmt.Printf("Philosopher %d finished dining\n", id)
        }(i)
    }
    wg.Wait()
}
