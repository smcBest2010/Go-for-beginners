package average

func AdditionMin(num1 int, num2 ...int) int {
	result := num1
	for i := 0; i < len(num2); i++ {
		if result > num2[i] {
			result = num2[i]
		}
	}
	return result
}
