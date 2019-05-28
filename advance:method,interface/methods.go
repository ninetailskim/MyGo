package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

type MyFloat float64
//belong to the vertex
func (v Vertex) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (f MyFloat) Abs() float64{
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}

func main(){
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	//what the fuck??? ok, i understand~
	fmt.Println(-math.Sqrt2)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}