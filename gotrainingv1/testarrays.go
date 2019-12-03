package main
import "fmt"

func main(){

	arr := [6] int {1,2,3,4,5,6}
	fmt.Println(arr[3])

	//untyped array

	var arr2 [2] string
	arr2[0] = "test"
	arr2[1] = "test1"

	fmt.Println(arr2[0])

	// default value for string array is nil and int array is 0
	var testint [2] int

	fmt.Println(testint[0])
	var testStr [2] string

	fmt.Println(testStr[0])

	//Array types are one dimentional but you can compose two D array
	var 2darray [3][4] int
	
	// for(int i =0; i < 3; i ++){
	// 	for (int j=0; j <4; j ++){
	// 		fmt.Println(i*j)
	// 	}
	// }

}