package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area() // Value Receiver - Copy해서 전달함
	println(area)

	/**
	메소드는 함수이지만 특정 struct를 지정해줘서 특정 struct(receiver) 전용 함수처럼 사용할 수 있는 것이다.
	Java는 클래스 내부에 상태 표시를 위한 필드와 상호 작용을 위한 메소드를 함께 정의하는데 Go에서는 struct 내부에 필드만 정의할 수 있다.
	*/

	area = rect.area2() // Pointer receiver - struct의 포인터(주소)만을 전달함
	println(area)
}
