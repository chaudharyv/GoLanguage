package main

import "fmt"

func sum(nums ... int) int{
	total :=0
	// not interested in indes so _ is used.
	for _, num := range nums {
		total +=num 
	}
	return total;
}

func main(){
	fmt.Println(sum(1,2,3))

	nums := []int{1,2,3,4,5,6}
	fmt.Println(sum(nums...))
}