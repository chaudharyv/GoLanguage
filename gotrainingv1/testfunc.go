package main
import "fmt"

//private method starts with small letter and public starts with Capital letters.
func plus( a int, b int) int{
return a+b;
}
func plusplus( a , b ,c int) int{
	return a+b+c;
}

//The (int,int) in this function signature shows that the function return two integers.
func test()(int,int){
	return 3,7
}

func addAndSubstract(a,b int)(int,int){
	return a+b, a-b;
}

func main(){
	res := plus(1,2)
	fmt.Println(res);
	res = plusplus(1,2,3)
	fmt.Println(res);
	a,b := test()
	fmt.Println("a,b" , a,b)
	a,b = addAndSubstract(2,2)
	fmt.Println("a,b" , a,b)
	// just interested in substracted result
	_,b = addAndSubstract(2,2)
	fmt.Println("substracted result" , b)
}
