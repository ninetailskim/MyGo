package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

func (v Vertex) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

//if there's no *, the scale will operate on a copy of v
func (v *Vertex) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}