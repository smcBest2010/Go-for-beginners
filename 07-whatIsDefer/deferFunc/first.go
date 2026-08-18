package deferFunc

import (
	"fmt"
)

func NameLength(name string) (int, string) {
	defer fmt.Println("im done") // @@ 여기서 디퍼는, 밑에있는 코드가 실패했든 성공했든 실행시킨다, 예외처리를 하기위해 defer가 존재한다고 보면 된다.
	return len(name), name
}
