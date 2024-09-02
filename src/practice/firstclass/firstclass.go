package main

func main() {
	add := func(i int, j int) int {
		return i + j
	} // function을 변수에 할당할 수 있음

	r1 := calc(add, 10, 20) // function을 다른 function의 파라미터로 사용할 수 있음
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

	r3 := calc2(func(x int, y int) int { return x * y }, 100, 200)
	println(r3)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

type calculator func(int, int) int

func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
