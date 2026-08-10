package printName

import (
	"strings"
)

func MyLengthUpper(name string) (string, int) {
	return strings.ToUpper(name), len(name)
}
