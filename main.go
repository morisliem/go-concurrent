package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	in := make(chan string)

	go task1(in)
	go task2(in)
	go task3(in)

	for i := 0; i < 3; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("time out")
			return
		case val := <-in:
			fmt.Println(val)
		}
	}

	close(in)

	finish := time.Since(start)
	fmt.Println(finish)
}

func task1(ch chan string) {
	fmt.Println("Starting task 1 ...")
	time.Sleep(4 * time.Second)
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
