package main

import "fmt"

func main(){
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = []int{2, 3, 5, 7, 11, 13}
	// here it's not like python ,the index must be non-negtive
	//ss = s[1:-2]
	//fmt.Println(s[-1])

	s = s[:0]
	printSlice(s)

	//what the fuck? s has been set by []??? 
	//how can it output [2,3,5,7]???
	//amazing~
	s = s[1:6]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = []int{2, 3, 5, 7, 11, 13}
}


func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}