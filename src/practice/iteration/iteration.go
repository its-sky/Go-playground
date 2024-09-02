package main

func main() {
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)

	for { // 무한 루프
		println("Infinite Loop")
	}

	names := []string{"홍길동", "이순신", "강감찬"}

	for idx, name := range names {
		println(idx, name)
	}

	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	println(a)

END:
	println("끝입니다")
}
