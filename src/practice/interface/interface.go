package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// Rect 타입에 대한 인터페이스 구현
func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	println(r.area())
	println(r.perimeter())

	println(c.area())
	println(c.perimeter())

	showArea(r, c)

	var x interface{}
	x = 1
	x = "Tom"

	EmptyInterface(x)

	var a interface{} = 1

	i := a       // dynamic type, 1
	j := a.(int) // int type, 1

	println(i) // 포인터 주소가 출력됨
	println(j)
}

// 이렇게 빈 인터페이스를 자주 볼 수 있는데 Go에서 모든 타입을 나타내기 위해 빈 인터페이스를 활용한다
// 어떤 타입이라도 담을 수 있는 컨테이너라고 보면 되며 Dynamic Type이라고 볼 수 있다. Java에서 Object와 같은 개념
func EmptyInterface(v interface{}) {
	fmt.Println(v)
}
