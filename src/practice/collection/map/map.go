package main

import "fmt"

func main() {
	var idMap map[int]string // map[key]value 로 선언함, nil map
	idMap = make(map[int]string)

	tickers := map[string]string{
		"GOOG": "Googile Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
	}

	var m map[int]string
	m = make(map[int]string)
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	println(str)

	noData := m[999] // 값이 없으면 nil 혹은 zero 반환
	println(noData)

	// 특정 key에 해당하는 값 삭제
	delete(m, 777)

	// key가 존재하는지 체크
	_, exists := tickers["MSFT"]
	if exists {
		println("MSFT ticker exists")
	}

	idMap[0] = "Apple"
	idMap[1] = "Banana"
	idMap[2] = "Pear"

	// Map에 존재하는 key-value 열거
	for key, val := range idMap {
		fmt.Println(key, val)
	}
}
