package main
import "fmt"

func main(){
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	// We don't need to declare that it is a variable here, := does the trick here
	i := "test"
	fmt.Println(i)
}