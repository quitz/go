package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
  
  //if the channel is not closed, line 19 will panic
  //error: fatal error: all goroutines are asleep - deadlock!
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}

}

//https://stackoverflow.com/questions/45020481/range-a-channel-finishes-with-deadlock
//https://stackoverflow.com/questions/52943450/go-routine-for-range-over-channels
