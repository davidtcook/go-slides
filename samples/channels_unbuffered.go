// channels_unbuffered
package main

import (
	"fmt"
	"time"
)

func channelWrite(ch chan<- int, value int) {
	fmt.Println("started channel write", value)
	ch <- value
	fmt.Println("finished channel write", value)
}

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go channelWrite(ch, i)
	}
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("received from channel", <-ch)
	}
	fmt.Println("main: exiting")
}
