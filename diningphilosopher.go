package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    forks  []*sync.Mutex
    states []int
)

func philosopher(i int) {
    for {
        // think
        states[i] = 0
        fmt.Println("Philosopher", i, "is thinking")
        time.Sleep(time.Second)

        // acquire left and right forks
        left := forks[i]
        right := forks[(i+1)%len(forks)]
        left.Lock()
        right.Lock()

        // eat
        states[i] = 1
        fmt.Println("Philosopher", i, "is eating")
        time.Sleep(time.Second)

        // release left and right forks
        right.Unlock()
        left.Unlock()

        // check if checkpoint reached
        if states[i] == 1 {
            break
        }
    }
}

func main() {
    // get the input for N
    fmt.Println("Enter the number of philosophers:")
    var N int
    fmt.Scanf("%d", &N)

    // create the forks
    forks = make([]*sync.Mutex, N)
    for i := 0; i < N; i++ {
        forks[i] = &sync.Mutex{}
    }

    // create the states
    states = make([]int, N)

    // start the philosophers
    for i := 0; i < N; i++ {
        go philosopher(i)
    }

    // for loop to keep the program running
    for {
    }
}
