package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// type parameter numbersToSum [][]int  (Slice of Slices)
	// Sample:
	// Invoke SumAll	->	SumAll([]int{1, 2}, []int{0, 9})
	// Internally here	->	numbersToSum := [][]int{{1, 2}, {0, 9}}
	lenSliceOfSlices := len(numbersToSum)
	sumsSlice := make([]int, lenSliceOfSlices)

	for i, numbers := range numbersToSum {
		sumsSlice[i] = Sum(numbers)
	}

	return sumsSlice
}
