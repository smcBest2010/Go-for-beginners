package main // Golang에서 시작되는건 main.go 파일에서 시작되는 main 패키지임을 의미함

import ( //
	"01-hello-world/something" // 같은 프로젝트 내에 있는 something 패키지를 import함
	"fmt"                      // fmt 패키지를 import함. fmt는 format의 약자로, 문자열을 출력하거나 포맷팅하는 기능을 제공함
)

func main() { // main 함수는 프로그램의 시작점임. main 패키지 안에 있어야 함
	fmt.Println("hello world") // fmt.Println() 함수는 문자열을 출력하는 함수임. 여기서는 "hello world"를 출력함
	something.SayHello()       // something 패키지 안에 있는 SayHello() 함수를 호출함. 이 함수는 something 패키지에서 정의되어 있음 하지만
}
