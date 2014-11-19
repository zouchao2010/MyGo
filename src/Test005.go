package main

import (
	"fmt"
)

const (
	KEYS string = "key"
	NULL byte   = 0x00
)

func byteArrayKey() []byte {
	keys := []byte(KEYS)
	length := 16 - len(keys)
	for i := 0; i < length; i++ {
		keys = append(keys, NULL)
	}
	return keys
}
func byteArray2longs(s []byte) []int {
	sLen := len(s)
	length := int((sLen + 3) / 4)
	for i := 0; i < sLen*4-length; i++ {
		s = append(s, NULL)
	}
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = int(s[i*4]) | int(s[i*4+1])<<8 | int(s[i*4+2])<<16 | int(s[i*4+3])<<24
	}
	return result
}
func btea(v []int, n int, k []int) {
	DELTA := 0x9e3779b9
	if n > 1 {
		z := v[n-1]
		q := 6 + 52/n
		_sum := 0

		for i := 0; i < q; i++ {
			_sum = (_sum + DELTA) & 0xffffffff
			e := (_sum >> 2) & 3
			for p := 0; p < n-1; p++ {
				y := v[p+1]
				v[p] = v[p] + (((z>>5^y<<2) + (y>>3^z<<4)) ^ ((_sum^y) + (k[(p&3)^e] ^ z)))
				v[p] &= 0xFFFFFFFF
				z = v[p]
			}
			y := v[0]
			v[n-1] = v[n-1] + (((z>>5^y<<2) + (y>>3^z<<4)) ^ ((_sum^y) + (k[(n-1&3)^e] ^ z)))
			v[n-1] &= 0xFFFFFFFF
			z = v[n-1]
		}
	} else if n < -1 {
		n = int(-n)
		q := 6 + 52/n
		_sum := q * DELTA
		y := v[0]

		for i := 0; i < q; i++ {
			e := ((_sum & 0xffffffff) >> 2) & 3
			for p := n - 1; p > 0; p-- {
				z := v[p-1]
				v[p] = v[p] - (((z >> 5) ^ (y << 2)) + ((y >> 3) ^ (z << 4)) ^ (_sum ^ y) + (k[(p&3)^e] ^ z))
				v[p] &= 0xFFFFFFFF
				y = v[p]
			}
			z := v[n-1]
			v[0] = v[0] - (((z >> 5) ^ (y << 2)) + ((y >> 3) ^ (z << 4)) ^ (_sum ^ y) + (k[(0&3)^e] ^ z))
			v[0] &= 0xFFFFFFFF
			y = v[0]
			_sum = _sum - DELTA*q
		}
	}
}
func longs2byteArray(s []int)([]byte){
	length :=len(s)
	fmt.Println("length...........",length)
	var result [length*4]byte
	fmt.Println("result...........",result)
	var result3 [10]byte
	fmt.Println("result3..........",result3)

//	for i:=0; i<length;i++{
//		v:=s[i]
//		result[i*4]=v&0xFF
//		result[i*4+1]=v>>8&0xFF
//		result[i*4+2]=v>>16&0xFF
//		result[i*4+3]=v>>24&0xFF
//	}
	var result2 []byte
//	for i:=0; i<len(result);i++{
//		if result[i]!=NULL{
//			append(result2,result[i])
//		}
//	}
	return result2

}
func encrypt(str string)(string) {
	byteKey := byteArrayKey()
	byteStr := []byte(str)
	v := byteArray2longs(byteStr)
	k := byteArray2longs(byteKey)
	var n int = len(v)
	//	fmt.Println(v)
	//	fmt.Println(k)
	fmt.Println("n..1.", n)
	fmt.Println("v..1.", v)
	fmt.Println("k..1.", k)
	btea(v, n, k)
	fmt.Println("n..2.", n)
	fmt.Println("v..2.", v)
	fmt.Println("k..2.", k)
	result := string(longs2byteArray(v))

	return result

}

func main() {
	sr := "i love you, ccc"
	enc := encrypt(sr)
	print(enc)
//	var dec = decrypt(enc)

}

func test(){
	z:=6513507
	y:=2032166262
	_sum:=2654435769
	k:=[]int{7955819, 0, 0, 0}
	p:=0
	e:=2
//	t:=( (((z & 0xffffffff)>>5)^(y<<2)) + (((y & 0xffffffff)>>3)^(z<<4))^(_sum^y) + (k[(p & 3)^e]^z) )
	t:=(((z>>5^y<<2) + (y>>3^z<<4)) ^ ((_sum^y) + (k[(p&3)^e] ^ z)))
	fmt.Println(t)

}
