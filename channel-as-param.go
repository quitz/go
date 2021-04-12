// chan <-          writing to channel (output channel)
// <- chan          reading from channel (input channel)
// chan             read from or write to channel (input/output channel)

package main
import "fmt"

func main() {
    ch := make(chan int, 3)
    processWrite(ch)
    processRead(ch)
}

//write to channel
func processWrite(ch chan<- int) {
    ch <- 2
    //s := <-ch
}

//read from channel
func processRead(ch <-chan int) {
    s := <-ch
    fmt.Println(s)
    //ch <- 2
}

