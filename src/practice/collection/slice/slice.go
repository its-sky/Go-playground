package main

import "fmt"

func main() {
	// Slice - 가변형 배열 (가변 크기)
	var slice []int
	slice = []int{1, 2, 3}
	slice[1] = 10
	fmt.Println(slice)
	// built-in println으로 하면 해시코드값이 나오고 fmt.Println을 하면 리스트 내부 값이 출력됨

	s := make([]int, 5, 10) // make(자료형, length(길이), capacity(배열 최대 길이))
	println(len(s), cap(s))

	var ss []int
	if ss == nil {
		println("Nil Slice - 슬라이스에 별도 길이와 용량을 지정하지 않으면 길이와 용량이 0임")
	}
	println(len(s), cap(s))

	s = []int{0, 1, 2, 3, 4, 5}
	s = s[2:5] // 부분 슬라이스 (마지막 숫자의 바로 직전 인덱스까지 slicing 됨, Python과 동일)
	fmt.Println(s)

	s = append(s, 2)
	s = append(s, 3, 4, 5)
	fmt.Println(s)
	/**
	  용량(capacity)이 남아있을 경우 length가 변경되며 데이터가 추가되고
	  용량 초과의 경우 현재 용량의 2배에 해당하는 새로운 array를 생성하고
	  기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당함
	*/

	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
	fmt.Println(sliceA) // 슬라이스에 슬라이스를 append할 수 있음. (...)은 elipsis라고 하며 슬라이스의 모든 요소를 표현함

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)
	println(len(target), cap(target))

	/**
	슬라이스는 내부적으로 사용하는 배열의 부분 영역인 세그먼트에 대한 메타 데이터를 가지고 있음
	슬라이스는 3가지 필드로 구성되어 있음.
	(내부적으로 사용하는 배열에 대한 포인터 정보, 세그먼트 길이, 세그먼트 최대 용량).
	[0 1 2 3 4 5] 라는 슬라이스를 [2 3 4]로 sub-slice했을 경우 len은 3, cap은 4가 된다.
	cap은 2부터 시작하여 가장 끝까지 길이를 재어 4를 갖게 되는 것이다.
	*/
}
