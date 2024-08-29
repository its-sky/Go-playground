package main

import "fmt"

func main() {
	// Int
	var a int = 3
	b := 3
	fmt.Println(a, b)
	b = 4
	fmt.Println(b)

	// Float
	var f float32 = 11.
	ff := 11.
	fmt.Println(f, ff)

	// 복수 개의 변수 한번에 선언
	var i, j, k int = 1, 2, 3
	fmt.Println(i, j, k)

	// String
	var s = "HI"
	fmt.Println(s)

	// 상수
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)
	fmt.Println(Visa, Master, Amex)

	const (
		Apple  = iota // 0, iota라는 identifier인데 0부터 시작하는 auto-increment임
		Grape         // 1
		Orange        // 2
	)
	fmt.Println(Apple, Grape, Orange)
}
