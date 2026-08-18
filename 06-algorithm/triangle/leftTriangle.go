package triangle

import (
	"fmt"
	"strings"
)

func Triangle01(num int) {

	for i := 1; i <= num; i++ {
		fmt.Println(strings.Repeat("*", i))
	}

}

func Triangle02(num int) {
	for i := num; i >= 0; i-- {
		fmt.Println(strings.Repeat("*", i))
	}
}
func Triangle03(num int) {
	for i := num; i >= 0; i-- {
		fmt.Println(strings.Repeat(" ", i) + strings.Repeat("*", num-i))
	}
}

func Triangle04(num int) {
	for i := num; i >= 0; i-- {
		fmt.Println(strings.Repeat(" ", i) + strings.Repeat("*", num-i) + strings.Repeat("*", (num-i)))

	}

	for i := 0; i <= num; i++ {
		fmt.Println(strings.Repeat(" ", i) + strings.Repeat("*", num-i) + strings.Repeat("*", (num-i)))
	}
}
