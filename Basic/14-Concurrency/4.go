package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500 ms"
			time.Sleep(time.Millisecond * 500)
			close(c1)
		}
	}()

	go func() {
		for {
			c2 <- "every two seconds"
			time.Sleep(time.Second * 2)
			close(c2)
		}
	}()

	for {
		//receiving on the channels in successive lines of code will not let us print concurrently, instead the code will receive them in order (c1 first and then c2, and then c1 and c2 and so on) which means that even though the goroutine involving c1 is completed in the background while the c2 goroutine is sleeping, it cannot be receieved and printed as c2 will get printed and only then will the code move on to c1 in the next iteration. This happens despite the fact that c1 receives in lesser time and would have been ready to print had we done our work concurrently. In this way, time is being wasted since the code is busy waiting for c2, during which time the execution for c1 can be commpleted and recieved. This problem can be addressed with the select statement, which can allow us to receive from whichever channel is available
		/*
			fmt.Println(<-c1)
			fmt.Println(<-c2)
		*/
		// --------- the select statement allows us to receive from whichever channel is ready
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case (<-c1): 
		}
	

		// case <-c1
	}
}
