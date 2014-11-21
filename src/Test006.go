package main

import (
	"fmt"
	"encoding/hex"
)

func main() {
	str :="ajfahfajhfk"
	strArr:=[]byte(str)
	fmt.Println(strArr)

	fmt.Println(string(strArr))

	fmt.Println(hex.DecodeString("3c497894518d601704864cbf98caeb9b"))

}
