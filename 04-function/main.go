package main

import (
	"chaos/lenAndUpper"
	"fmt"
)

func main() {
	totalLen, upperName := lenAndUpper.LaU("nico")
	fmt.Printf("%d개의 길이를 가지고 있고, 대문자로 하면 %s\n", totalLen, upperName) // C와 같이 Printf를 활용하여 문자열 포매팅이 가능하다. 하지만 줄바꿈은 안된다.

	totalLen02, _ := lenAndUpper.LaU("hello")
	fmt.Printf("%d 생락하려면 언더바를 활용하자,", totalLen02) // 언더바를 사용하면 value값을 무시하게 된다.

}
