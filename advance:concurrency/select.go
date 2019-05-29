package main

import "fmt"

func fibonacci(c, quit chan int){
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//1. fibonacci to resquest c and quit
//2. c is in go routines and +++, quie is wait after the c
//3. output all c , then quit case , then return

func main(){
	c := make(chan int)
	quit := make(chan int)
	go func(){
		for i := 0; i < 10;i ++{
			fmt.Println(<-c)
			//quit <- i
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}