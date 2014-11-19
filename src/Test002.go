package main

import "fmt"

func main() {
	fmt.Println("Hello")
	x := 0
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	var y string = "asldfjl"
	fmt.Println(y)
	fmt.Println(len(y))

	fmt.Println(3/2)
	fmt.Println(3/1)
	fmt.Println(3%2)


}
