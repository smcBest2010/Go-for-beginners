package something

import "fmt"

func sayBye() { // 함수를 소문자로 시작하면 외부 패키지에서 접근 불가능함. 즉, private 함수임
	fmt.Println("Bye")
}

func SayHello() { // 함수를 대문자로 시작하면 외부 패키지에서 접근 가능함. 즉, public 함수임
	fmt.Println("Hello")
}
