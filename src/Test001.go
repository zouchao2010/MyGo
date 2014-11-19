package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var i int
	i = 100
	fmt.Printf("Value is: %v\n", i)
	fmt.Println(i + 11)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)

	fmt.Println(v)

	fmt.Println(e, f, g)
}

const (
	a = iota
	b
	c
)
const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= i
)
const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
const (
	e, f, g = iota, iota, iota //e=0,f=0,g=0 iota在同一行值相同
)
