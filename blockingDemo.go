package main

/*
	Demo for how sending and receiving messages across channels is a blocking operation
*/

import "fmt"

func main() {
	/*
		c:= make(chan string) //Channel of strings
		c <- "hello" //Sending hello across the channel

		msg := <- c //Trying to receive msg from the channel
		fmt.Println(msg) //Output msg to STDIO

		This doesn't work coz the send will block until something is ready to receive
		The code never progresses to the receive line because we're blocked at send

		To make this work, we'd need to receive in a separate go routine OR ALTERNATIVELY
		We can make a buffered channel by giving a capacity when we make the channel
	*/

	//You can fill up a buffered channel without a corresponding receiver and it won't block until the channel is full
	c := make(chan string, 2)
	c <- "hello dave"
	c <- "you're kickass"

	msg := <- c
	fmt.Println(msg)

	msg2 := <- c
	fmt.Println(msg2)

}
