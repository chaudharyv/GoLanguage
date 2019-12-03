package main
import "fmt"

func intSeq() func() int {
	i := 0
	return func() int{
		i++
		return i
	}
}

func main(){
	res := intSeq()
	fmt.Println(res())
	fmt.Println(res())
	fmt.Println(res())
}