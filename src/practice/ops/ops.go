package main

import "fmt"

func main() {
	a := 3
	b := 5
	c := (a + b) / 2

	fmt.Println(a, b, c)
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	i := 1
	i++
	fmt.Println(i) // i++을 출력하는 것이 안됨

	fmt.Println(true && false)
	fmt.Println(true || false)

	fmt.Println(2 & 3 << 2) // 10 & 11 = 10 << 2  1000 = 8

	var k int = 10
	var p = &k
	fmt.Println(*p)
}
