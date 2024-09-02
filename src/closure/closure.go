package main

func main() {
	next := nextValue()

	println(next())
	println(next())
	println(next()) // 3

	newNext := nextValue()
	println(newNext())
	println(newNext())
	println(newNext()) // 3
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*
Go 함수는 Closure로써 사용될 수 있는데 함수 바깥에 있는 변수를 참조하는 함수 값을 일컫는다.
nextValue() 함수는 int를 리턴하는 익명함수를 리턴하는 함수이다.
Go 언어에서느 함수가 다른 함수로부터 리턴되는 리턴값으로 사용될 수 있는 일급함수의 성격을 지니고 있다.
*/
