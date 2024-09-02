package main

func main() {
	var k int = 0

	if k == 1 {
		println("k is one")
	} else if k == 2 {
		println("k is two")
	} else {
		println("k is neither one nor two")
	}

	var i int = 1
	var maximum int = 5
	if val := i * 2; val < maximum {
		println(val)
	}
	//++val - if scope를 벗어나서 에러가 뜸

	var name string
	var category = 1

	switch category {
	case 1:
		name = "A"
	case 2:
		name = "B"
	case 3:
		name = "C"
	case 4, 5:
		name = "D"
	default:
		name = "OTHER"
	}
	println(name)

	switch x := category << 2; x - 1 {
	case 3:
		name = "is 3"
	default:
		name = "just default"
	}
	println(name) // is 3

	grade(70)
}

func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
		fallthrough
	default:
		println("No..")
	}
}
