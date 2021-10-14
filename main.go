package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Second

func main() {
	ctx := context.Background()
	start := time.Now()
	in := make(chan string)

	ctx, cancel := context.WithTimeout(ctx, 2000000000)
	defer cancel()

	go task1(in)
	go task2(in)
	go task3(in)

	for i := 0; i < 3; i++ {
		select {
		case val := <-in:
			fmt.Println(val)
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}

	close(in)

	finish := time.Since(start)
	fmt.Println(finish)
}

func task1(ch chan string) {
	fmt.Println("Starting task 1 ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Task 1 finished")
	ch <- "Task 1"
}

func task2(ch chan string) {
	fmt.Println("Starting task 2 ...")
	time.Sleep(2 * time.Second)
	fmt.Println("Task 2 finished")
	ch <- "Task 2"
}

func task3(ch chan string) {
	fmt.Println("Starting task 3 ...")
	time.Sleep(1 * time.Second)
	fmt.Println("Task 3 finished")
	ch <- "Task 3"
}
