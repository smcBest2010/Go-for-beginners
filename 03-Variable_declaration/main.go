package main

import (
	"fmt"
	"orange/printName"
)

func parameter(name string, num int) (string, int) { //@@ 파라미터 오른쪽의 괄호는 리턴의 타입을 또 정의해야하므로 써주자
	return name, num
}
func parameter02(alpha string) string {
	return alpha
}
func main() {
	/*
		name := "홍길동"     // 해당 문구로 변수를 정의 할 수 있음 str타입으로 정해지게된다
		name = "홍진호"      // 변수 값을 바꿀수 있다.
		fmt.Println(name) // 프린트라인 변수 포매팅방법
	*/
	printName.PrintMyname(parameter("홍길동 ", 10))
	fmt.Println(printName.MyLengthUpper(parameter02("hello")))
}
