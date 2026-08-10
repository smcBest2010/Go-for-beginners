package lenAndUpper

import (
	"strings"
)

func LaU(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
