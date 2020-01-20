package main

/*
	Go's select statement allows us to avoid blocking code.
	Let a line of code continue executing if it can
*/

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	//In the main go routine ...
	for {
		/*
			This is blocking code that causes us to receive one value after the other
			We're blocking while waiting for the slow one

			fmt.Println(<- c1)
			fmt.Println(<- c2)

			Use a select to receive from whichever channel is ready. You'll receive quickly from channel 1
		*/
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}


}
