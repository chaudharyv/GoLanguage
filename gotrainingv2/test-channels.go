package main
import "fmt"

func main(){
	// create a new channel with make
	messages := make(chan string)

	// send a value into a channel using the channel
	go func() { messages <- "ping"}()
	// he <- channel syntax recieves a value from the channel.
	
	msg := <- messages

	fmt.Println(msg)
}
