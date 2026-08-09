package operator // 패키지명을

import (
	"fmt"
)

func Double(x int) { // 2배로 불려주는 함수 퍼블릭임 매개변수
	fmt.Println(x * 2)
}
