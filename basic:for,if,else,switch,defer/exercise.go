package main

import (
	"fmt"
)

func getResult(x float64) float64{
	z := 1.0
	time := 0
	for ; time < 10; time ++{
		z -= (z * z - x) / (2 * z)
		fmt.Printf("time %d : z: %f\n", time, z) 
	}
	return z
}

func main(){
	fmt.Println("Final: %f\n", getResult(25) )
}