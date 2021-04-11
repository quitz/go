package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// Writer (launched in own goroutine)
	go func() {
		time.Sleep(time.Duration(500) * time.Millisecond)
		fmt.Println("Writing value 1515...")
		// This send statement is a blocking operation for this thread
		ch <- 1515
		// Maybe the following won't have time to execute : never mind
		fmt.Println("Value 1515 written.")
	}()

	// Reader: in main thread
	time.Sleep(time.Duration(1000) * time.Millisecond)
	fmt.Println("Reading first value (if exists)...")
	// This receive operation is blocking,
	// and there is a value ready for reading
	select {
	case x, ok := <-ch:
		if ok {
			fmt.Printf("Value %d was read.\n", x)
		} else {
			fmt.Println("Channel closed!")
		}
	default:
		fmt.Println("No value ready, moving on.")
	}

	fmt.Println("Reading second value (if exists)...")
	// This receive operation is blocking,
	// and there is no value to read : BIG DEADLOCK HAZARD
	select {
	case x, ok := <-ch:
		if ok {
			fmt.Printf("Value %d was read.\n", x)
		} else {
			fmt.Println("Channel closed!")
		}
	default:
		fmt.Println("No value ready, moving on.")
	}

	fmt.Println("The end")
}


//https://stackoverflow.com/questions/3398490/checking-if-a-channel-has-a-ready-to-read-value-using-go
