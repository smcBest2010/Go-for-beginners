package main // 패키지명을 main으로 정의하여 시작점을 컴파일러에게 남깁니다.

import (
	"apple/operator" // go.mod 기준으로 operator디렉토리를 찾아 함수를 불러옵니다.
)

func main() { // 엔트리 포인트 시작점입니다.
	operator.Double(2)
}
