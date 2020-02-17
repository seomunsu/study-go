package main

import "fmt"

func main() {
	// 중첩 구조체
	// ex1
	car1 := Car1{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},
	}

	// 내부 spec 구조체 값 출력
	fmt.Println("ex1 :", car1.name)
	fmt.Println("ex1 :", car1.color)
	fmt.Println("ex1 :", car1.company)
	fmt.Printf("ex1 : %#v\n", car1.detail)

	// ex2
	// 내부 spec 구조체 필드 값 출력
	fmt.Println("ex2 :", car1.detail.length)
	fmt.Println("ex2 :", car1.detail.height)
	fmt.Println("ex2 :", car1.detail.width)
}

// 첫 글자 대문자로 선언 -> public
type Car1 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

// 첫 글자 소문자로 선언 -> private
type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}
