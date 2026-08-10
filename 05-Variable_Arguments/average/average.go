package average

// 가변인자를 이용한 함수지만, 1개 이상의 값을 파라미터에 넣어야 작동된다. 해당 num2파라미터는 슬라이스(리스트) 형태
func Avg(num1 float64, num2 ...float64) float64 {
	count := num1
	for i := 0; i < len(num2); i++ {
		count += num2[i]
	}
	return count / float64(len(num2)+1)
}
