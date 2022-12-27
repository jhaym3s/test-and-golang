package slice

func Sum(numbers [5]int) int {
	initialValue := 0
	for i := 0; i < 5; i++ {
		initialValue = initialValue + numbers[i]
	}
	return initialValue
}