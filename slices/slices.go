package slice

// func Sum(numbers [5]int) int {
// 	initialValue := 0
// 	for i := 0; i < 5; i++ {
// 		initialValue = initialValue + numbers[i]
// 	}
// 	return initialValue
// }

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	capacityOfReturnedSlice := len(numbersToSum)
	//use "len" to get the capacity for the slice to be returned.
	sums := make([]int,capacityOfReturnedSlice)
	//above you are making a new slide with with the capacity gotten before

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}