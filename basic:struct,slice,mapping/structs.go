package main

import "fmt"

type Vertex struct{
	X int
	Y int
}


var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}// Y default seted by ZeroValue
	v3 = Vertex{}
	pp = &Vertex{3, 4}
)
func main(){
	v := Vertex{1,2}
	p := &v
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(p.X)

	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, v2, v3, pp)
}