package main

import "fmt"

func main(){
	// make is like new operator and creating an array
	s := make([]string, 3)
	fmt.Println("emp : ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("emp s: ", s)
	fmt.Println("len s: ", len(s))

	c := make([]string, 3)
	c = s
	fmt.Println("emp c: ", c)
	fmt.Println("len c: ", len(c))

	//slice starting from 3
	l := c[:3]
	fmt.Println("sl2: ", l)

	//slice from 2 till end
	l = c[2:]
	fmt.Println("sl2: ", l)
}