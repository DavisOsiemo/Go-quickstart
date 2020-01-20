package main

import (
	"fmt"
)

/*
	Where you have a queue of work to be done and multiple concurrent workers pulling items off the queue

	Use fibonacci algorithm to reinforce this concept
*/

func main() {
	//Create both channels. Make them buffered channels
	jobs	:= make(chan int, 100)
	results	:= make(chan int, 100)

	//Create a worker as a concurrent go routine
	go worker(jobs, results)

	//Add more workers to pull items off the jobs queue and putting the result into the results queue/channel
	//You'll end up using more CPU resources, which means the work is executed faster. (A good thing since we're taking advantage of the multi-core processor)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	//Fill out the jobs channel with 100 numbers (Iterate no.s 0 - 99).
	//Since it's buffered, no blocking, so no worries
	for i := 0 ; i <= 100; i ++ {
		jobs <- i
	}
	close(jobs) //Close channel coz we're done putting stuff into that channel

	/*
		Once the worker is done adding the values, it begins pulling off one number at a time from the channel and calculating the Fib
		It then puts back the results into the results channel
	*/

	for j:= 0; j <= 100; j++ {
		fmt.Println(<-results)
	}
	fmt.Scanln()
}

/*
	Takes in 2 channels
	Jobs channel	= jobs to do (only ever receive on the jobs channel)
	Results channel	= results of work done (only ever send on the results channel)

	- Specifying which channels can only ever receive or send reduces the risk of bugs
	- When you try to send on the jobs channel you get a compile time error
*/

func worker (jobs <-chan int, results chan<- int) {
	//Jobs is a queue of numbers. Use a range feature to consume items from this queue
	for n:= range jobs {
		results <- fib(n) //Receive n from the channel, calculate the n'th Fibonacci number and send it on the results channel
	}

}

func fib (n int) int {
	if n <= 1 {
		return n
	}

	return fib(n - 1) + fib(n - 2)
}
