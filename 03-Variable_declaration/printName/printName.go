package printName

import (
	"fmt"
	"strings" // strings 모듈이란 문자열을 마음대로 가꿀수있도록
)

func PrintMyname(num1 string, num2 int) {
	fmt.Println(strings.Repeat(num1, num2)) // 스트링즈 모듈의 repeat기능을 활용
}
