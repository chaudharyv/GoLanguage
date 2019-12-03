package main

import "fmt"
import "math"

type geometry interface{
	area() float64
	perim() float64
}

type rect struct{
	width, height float64
}

type circle struct{
	radius float64
}

//To implement an interface in Go, we just need to implement the methods

func (r rect) area() float64{
	return r.width * r.height
}

func (r rect) perimeter() float64{
	return 2*r.width + 2*r.height
}

func (c circle) area() float64{
   return math.Sqrt(c.radius)
}

func main(){
	r := rect{width: 10, height :15}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
	c := circle{radius: 5}
	fmt.Println(c.area())
}