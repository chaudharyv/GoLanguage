package main
import "fmt"

type rect struct{
	width, height int
}

//this area method has a reciever type of *rect
// There (r *rect) says that there is a strcture called rect to which i am tying this method
func (r *rect) area() int{
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver

func (r rect) perimeter() int{
	return 2*r.width + 2*r.height
}

func main(){
	r := rect{width:10, height:10}
	area := r.area()
	fmt.Println("area is ", area)
	perimeter := r.perimeter()
	fmt.Println("perimeter is ", perimeter)

	// Go automatically handles converstion between values
	rp := &r
	fmt.Println("area is ", rp.area())
	fmt.Println("perimeter is ", rp.perimeter())
}