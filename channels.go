package main

/*
	Continue from concurrency.go

	Channels - a way for go routines to communicate with each other
			 - It's a pipe through which you can send/receive messages
			 - You can accept a channel as an arg. to a function
*/
import (
	"fmt"
	"time"
)
func main() {
	c := make(chan string)
	go count1("sheep", c)

	/*
		msg := <- c //Receive a message from the channel
					//Note that sending and receiving are blocking operations
					//When you want to receive a value/message you have to wait for there to be a value to receive
					//The blocking nature of channels allows us to synchronize go routines
		fmt.Println(msg)
	*/

	//Receive multiple messages using a for loop

	/*
		for {
			msg, open := <- c //You can receive more than one message
							  //The second message being one that informs you whether the channel is STILL open

			if !open {
				break //If channel is closed, get out of the for loop
			}
			println(msg)
		}
	*/

	//Better way to do the above is to iterate over the range of the channel
	for msg := range c {
		//Keep receiving messages and put them in the msg variable
		//No need to check manually is the channel is closed
		fmt.Println(msg)
	}

}

func count1(thing string, c chan string) {
	for i:=1; i<=5; i++ {
		c <- thing //Sending a message into a channel
				   //Also blocking. It will have to wait until a receive is ready to receive a message
		time.Sleep(time.Millisecond * 500)
	}
	/*
		This leads to a deadlock coz the count function completes sending a message while the main go routine still expects a message on the channel
		Program could potentially wait forever without terminating, Go detects this at runtime; Can't do so at compile time yet
		But when it happens it can notice that go routines aren't making any progress

		Therefore close the channel once you're done
		Only the sender should close the channel, not the receiver!!!!!
	*/

	close(c)
}
