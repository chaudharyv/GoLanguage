package main

import "os"

func main(){
	// uses panic to handle the error.
	panic("a problem")
    // if file doesn't exist then error
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}