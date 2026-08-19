package main

import (
	"fmt"
)

func main() {
	a := 12
	b := &a
	fmt.Println(&a, b)
	/*
	   윈도우에서는 exe파일로 볼수있는데 exe파일이 실행되면 프로그램이 램에서 프로세스 상태로 바뀜, 이제 그 프로세스중인 변수, 메서드의 값만 가져오는게 아니라 참조만하는게 바로 포인터다
	*/
}
