package main

import (
	"fmt"
	"time"
	"sync"
)


func main() {
	now := time.Now()
 
	d := 200
	for i:=0;i<d;i++{
    //this could be replaced by any io operations
		time.Sleep(time.Millisecond * 1)
	}
	
	now1 := time.Now()
	fmt.Println("elapsed time with regular for loop", time.Since(now))
	
	var wg sync.WaitGroup
	wg.Add(d)
	for i:=0;i<d;i++{
		go func(i int) {
		   defer wg.Done()
           time.Sleep(time.Millisecond * 1)
		}(i)
	}

	wg.Wait()
	fmt.Println("elapsed time with goroutines", time.Since(now1))
}


//Output
//elapsed time with regular for loop 227.69952ms
//elapsed time with goroutines 2.291133ms
