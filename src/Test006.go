package main

import "fmt"

func main() {

	arr := make([]int,10)
	append(arr,10)
	append(arr,10)
	append(arr,10)
	fmt.Println(arr)
//	var result [int(len(arr) * 4)]byte
//	fmt.Println(result)
//	longs2byteArray(arr)
}

//func longs2byteArray(s [10]int) []byte {
//	length := len(s)
//	length = 0
//	fmt.Println("length...........", length)
////	var result [int(length * 4)]byte
//	var result []byte
//	for i := 0; i < len(s); i++ {
//		fmt.Println("i...........", i)
//		append(result, 0x00)
//	}
//	fmt.Println("result...........", result)
//	var result3 [10]byte
//	fmt.Println("result3..........", result3)
//	var result2 []byte
//	return result2
//}
