package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	time.Sleep(1 * time.Second)
	fmt.Println("Hello world goroutine wait completed")
	for i := 0; i < 100; i++ {
		done <- true
		fmt.Println("Updated true to go channel ", i)
		// done <- false
		// fmt.Println("Updated false to go channel", i)
	}
	close(done)
}
func main() {
	done := make(chan bool, 1)
	go hello(done)
	for i := 0; i < 100; i++ {
		select {
		case donetest := <-done:
			fmt.Println("recieved from go channel ", donetest)
		}
	}
	// donetest := <-done
	// fmt.Println("recieved from go channel ", donetest)
	// donetest = <-done
	// fmt.Println("recieved from go channel ", donetest)
	fmt.Println("main function")
}
