package main

import "fmt"

type I interface{
	M() //the type? what the fuck
}

type T struct{
	S string
}

func (t T) M(){
	fmt.Println(t.S)
}

func main(){
	var i I = T{"Hello"}
	i.M()
}

//what the fucking difference with last one?