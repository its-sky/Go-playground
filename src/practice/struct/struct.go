package main

import "fmt"

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func main() {
	p := person{}

	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	p1 := person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	p3 := new(person) // new를 사용할 시 모든 필드를 zero value로 초기화하고 객체의 포인터를 리턴한다.
	p3.name = "Lee"

	/**
	Go 에서 기본적으로 struct는 mutable하다.
	필드 값이 변화할 경우 별도 새 개체를 만들지 않고 해당 개체의 메모리에서 직접 변경한다.
	하지만, struct 개체를 다른 함수의 파라미터로 넘긴다면 Pass by value에 따라 복사해서 전달한다.
	Pass by reference로 전달하고 싶을 때는 포인터를 전달하면 된다.
	*/
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	dic := newDict()
	dic.data[1] = "A"
	fmt.Println(dic.data[1])
}

// 생성자 함수 (constructor function)
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}
