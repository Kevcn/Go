package main

import (
	"fmt"
	"time"
)

func main() {

	// fmt.Println("------ Basic channel example ------")
	// c := make(chan string)
	// // its like putting this onto a thread to execute async
	// go count("sheep", c)

	// // consume all the message from channel c, automatically checks for closed
	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	fmt.Println("------ Multiple channel example ------")
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}

}

func count(input string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- input
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

func worker(jobs <-chan int, results chan<- int) {
	// pulls a job off the jobs channel and calculates the fib number, place it into the results channel
	for job := range jobs {
		results <- fib(job)
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
