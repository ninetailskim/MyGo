package main

import "fmt"

type Vertex struct{
	X int
	Y int
	z bool
	M float64
	//how can i set some dafault value by myself
	//N :"Nice"
}

// so if you want to only set part params, just follow the format "{X:xxx, Y: vvv}"

var (
	v1 = Vertex{X: 3, Y: 4}
	v2 = Vertex{X: 1}// Y default seted by ZeroValue
	v3 = Vertex{}
	v4 = Vertex{X: 100, z: true}
)
func main(){
	fmt.Println(v1, v2, v3, v4)
}