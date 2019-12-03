package main

import "fmt"

func f (from string){
	for i:=0; i<3; i++{
		fmt.Println(from, ":", i)
	}
}

func main(){
	// Direct method call
	f("direct")

	go f ("goroutine")

	go func(msg string){
		fmt.Println(msg)
	}("goroutineanonymous")

	fmt.Scanln()
	fmt.Println("done")

}