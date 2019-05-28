package main

import "fmt"

func main(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	//if you change the slice, the original data will be modified at the same time

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{{2, true},{3, false},{5, true},{7, true},{11, false},{13, true}}

	fmt.Println(s)

	p := []int{2,3,5,7,11,13}

	fmt.Println(&q, &p)
	fmt.Println(&q[0], &p[0])

	fmt.Printf("Type: %T\n", &p)
	fmt.Printf("Type: %T\n", p)

	pp := &p
	fmt.Printf("Type: %T\n", &pp)
	fmt.Println(&pp)
	fmt.Printf("Type: %T\n", pp)
	fmt.Println(pp)
	//how can i find the address of p, or does it equal to p[0]?

	p[0] = 100
	fmt.Println(q,p)
}