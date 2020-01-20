package main

//Concurrency in Go - Jake Wright


import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		Synchronous code(Counts sheep forever and never gets to fish)
		count("sheep")
		count("fish")
	*/

	/*
		With a go routine program doesnt wait to go to the next task. Run code concurrently
		Main function is also a go routine
		You can write as many go routines as you want but it doesn't guarantee your code will be infinitely fast. That's coz you're constrained by CPU cores
	*/

	/*
		go count("sheep")
		go count("fish")

		This code will cause the program to exit
		That's coz the when the main go routine finishes, all other routines finish regardless of whether they'd finished executing or not
		Initially the main couldn't finish because it had an infinite loop "count("fish")" so it dint exit
		But now we pushed the function count("fish") in its own go routine i.e go  count("fish") and nothing follows it, so the main go routine exits

		You could output for 2s with time.Sleep(time.Second * 2) you'd see the output of the go routines for 2 seconds
		Or you could use a blocking call that wait for user input i.e fmt.Scanln

	*/
	/*go count("sheep")
	go count("fish")*/

	//time.Sleep(time.Second * 5)
	//fmt.Scanln() - Run until you have user input (Not very useful in reality coz it requires manual user input)

	/*
		As a solution to the blocking code above, use a wait group as follows
	*/
	var wg sync.WaitGroup //It's just a counter, nothing scary
	wg.Add(1) //Means that I have only one go routine to wait for

	//When the go routine finishes (loop), decrement the counter as follows
	//Use anonymous function (no name) as follows
	//This syntax creates a function and immediately calls it

	go func() {
		count("sheep")
		wg.Done() //Done decrements counter by 1
	}()

	wg.Wait()

}

func count (thing string) {
	for i:= 1; i<=5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
/*
	Once this loop completes, decrement the counter
*/

/*
	Infinite loop implementation

	func count(thing string) {
		for i:= 1; true; i++ {
			fmt.Println(i, thing)
			time.Sleep(time.Millisecond * 500)
		}
}*/