package main
import "fmt"
import "math"

const s string = "constant"

func main(){
	fmt.Println(s)
	
	const n = 500
	 
	fmt.Println(math.Exp2(n))
	fmt.Println("a + b = ", 2+3)
	fmt.Println("Vipul " + "chaudhary")
}