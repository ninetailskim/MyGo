package main

import "fmt"

var i, j int = 1, 2

func main() {
	//can ignore the type of var
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}