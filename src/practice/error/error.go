package main

import (
	"log"
	"os"
)

type error interface {
	Error() string
}

func main() {
	f, err := os.Open("/Users/sinmincheol/Documents/github.png")
	if err != nil {
		log.Fatal(err.Error())
	}
	println(f.Name())

	// 타입 별로 다르게 에러를 핸들링할 수 있음
	/*_, err := otherFunc()
	switch err.(type) {
	default:
		println("No Error")
	case MyError:
		log.Println("Log My Error")
	case error:
		log.Fatal(err.Error())
	}*/
}
