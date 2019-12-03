package main

import "fmt"

//pass by value
func zeroval(ival int){
	ival = 0
}
//pass by reference
func zeropointer(iptr *int){
	*iptr = 0
}

func main(){
	i:= 1
	zeroval(i)
	fmt.Println("i", i);
	zeropointer(&i)
	fmt.Println("i", i);
}