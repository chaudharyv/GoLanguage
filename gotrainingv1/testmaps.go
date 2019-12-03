package main

import("fmt")

func main(){
	// built in method to create a map.
	m := make(map[string] int)
	m["k1"] = 1
	m["k2"] = 2
	v1 := m["k1"]
	fmt.Println("value: ", v1)
	fmt.Println("len: ", len(m))
	
	// Deleting the entry for the key k2
	delete(m, "k2")
	fmt.Println("len: ", len(m))

	// The optional second return value when getting a value
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	// initialize and declare map

	n := map[string]int{"foo" :1, "bar":2}
	fmt.Println("new map = " ,n)





}