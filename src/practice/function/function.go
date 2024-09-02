package main

func main() {
	msg := "Hello"
	say(msg)           // Pass by value
	sayReference(&msg) // Pass by reference
	say(msg)

	sayWithVaradic("a", "b", "ccccc", "dddd")

	println(sum(1, 2, 3, 4))
	println(sums(1, 2, 3, 4))
	println(sumsWithNamedReturnParam(1, 2, 3))
}

func say(msg string) {
	println(msg)
}

func sayReference(msg *string) {
	println(*msg)
	*msg = "Changed"
}

func sayWithVaradic(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sums(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
} // Multi return parameter

func sumsWithNamedReturnParam(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return // return은 항상 명시적으로 써줘야 함
}
