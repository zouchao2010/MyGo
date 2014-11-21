package main

import (
	"fmt"
	"encoding/hex"
)

const (
	KEYS string = "key"
	NULL byte   = 0x00
	StrTable string = "0123456789abcdef"
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
				v[p] = v[p] - (((z>>5^y<<2) + (y>>3^z<<4)) ^ ((_sum^y) + (k[(p&3)^e] ^ z)))
				v[p] &= 0xFFFFFFFF
				y = v[p]
			}
			z := v[n-1]
			v[0] = v[0] - (((z>>5^y<<2) + (y>>3^z<<4)) ^ ((_sum^y) + (k[(0&3)^e] ^ z)))
			v[0] &= 0xFFFFFFFF
			y = v[0]
			_sum = _sum - DELTA
		}
	}
}
func longs2byteArray(s []int)([]byte){
	var result []byte
	length :=len(s)
	for i:=0; i<length;i++{
		v:=s[i]
		result=append(result,byte(v&0xFF))
		result=append(result,byte(v>>8&0xFF))
		result=append(result,byte(v>>16&0xFF))
		result=append(result,byte(v>>24&0xFF))
	}
	return result

//	var result2 []byte
//	length2 :=len(result)
//	for i:=0; i<length2;i++{
//		if result[i]!=NULL{
//			result2=append(result2,result[i])
//		}
//	}
//	fmt.Println("longs2byteArray..........result2.....",result2)
//	return result2
}
func byteArray2Hex(s []byte)(string){
	result := ""
	for i:=0; i<len(s);i++{
		result += string(StrTable[s[i]>>4]) + string(StrTable[s[i]&0xF])
	}
	return result
}
func hex2ByteArray(s string)([]byte){
	result,_:=hex.DecodeString(s)
	return result
}
func encrypt(str string)(string) {
	byteKey := byteArrayKey()
	byteStr := []byte(str)
	v := byteArray2longs(byteStr)
	k := byteArray2longs(byteKey)
	var n int = len(v)
	btea(v, n, k)
	result := longs2byteArray(v)
	result2 := byteArray2Hex(result)
	return result2
}
func decrypt(str string)(string) {
	byteStr:=hex2ByteArray(str)
	byteKey := byteArrayKey()
	v := byteArray2longs(byteStr)
	k := byteArray2longs(byteKey)
	n := len(v)

	btea(v, -n, k)
	result := longs2byteArray(v)
	result2 := string(result)
	return result2

}

func main() {
	for i:=0; i<10000000;i++{
		sr := "i love you, ccc"
		enc := encrypt(sr)
		decrypt(enc)
		if (i%10000==0){
			fmt.Println(i)
		}
	}
}
//func main() {
//	sr := "i love you, ccc"
//	enc := encrypt(sr)
//	fmt.Println("sr...", sr)
//	fmt.Println("enc...", enc)
//	var dec = decrypt(enc)
//	fmt.Println("dec...", dec)
//
//}

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

func test2(){
	z:=3209463300
	y:=2490911036
	_sum:=50434279611
	k:=[]int{7955819, 0, 0, 0}
	p:=3
	e:=2
	t:= (((z>>5^y<<2) + (y>>3^z<<4)) ^ ((_sum^y) + (k[(p&3)^e] ^ z)))
//	t:=( (((z & 0xffffffff)>>5)^(y<<2)) + (((y & 0xffffffff)>>3)^(z<<4))^(_sum^y) + (k[(p & 3)^e]^z) )
//	t:=( (((z)>>5)^(y<<2)) + (((y )>>3)^(z<<4))^(_sum^y) + (k[(p & 3)^e]^z) )
	fmt.Println(t)
	fmt.Println((2615921304-25028488620)& 0xFFFFFFFF)
}
