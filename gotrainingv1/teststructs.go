package main

import "fmt"

type person struct{
	name string
	age int
}

func main(){
	fmt.Println(person{"BOB", 20})
	fmt.Println(person{name:"ALICE", age:20})
}