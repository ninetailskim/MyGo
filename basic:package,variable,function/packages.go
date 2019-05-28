package main

import (
	"fmt"
	"math/rand"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main(){
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Now you have %g problem.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	var i int
	fmt.Println(i,c, python, java)
	fmt.Println(split(17))
}