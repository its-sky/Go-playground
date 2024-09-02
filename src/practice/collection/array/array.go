package main

func main() {
	// Array
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	for _, item := range a {
		println(item)
	}

	//var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3, 4}

	for _, item := range a3 {
		println(item)
	}

	var multiArray [3][4][5]int
	multiArray[0][1][2] = 10

	var arr = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println(arr[1][2])
}
