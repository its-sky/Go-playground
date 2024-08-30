package main

import "fmt"

func main() {
	var boolean bool = false

	var str string = "str" // immutable

	var a int8 = 127          // -128 ~ 127, 8비트. 1비트는 sign이라 2^7
	var b int16 = 32767       // -32768 ~ 32767
	var c int32 = -2147483648 // -2147483648 ~ 2147483647

	rawLiteral := `아리랑\n
아리랑\n
아라리요`

	interLiteral := "아리랑아리랑\n아리리요"

	fmt.Println(boolean, str, a, b, c)
	fmt.Println(rawLiteral)
	fmt.Println(interLiteral)

	// Type Conversion
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	fmt.Println(u, f)

	newStr := "ABC"
	bytes := []byte(newStr)
	str2 := string(bytes)
	fmt.Println(newStr, bytes, str2)
}
