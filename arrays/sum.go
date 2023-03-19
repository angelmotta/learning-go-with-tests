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
	lenSlice := len(numbersToSum)
	sumsResult := make([]int, lenSlice)

	for i, numbers := range numbersToSum {
		sumsResult[i] = Sum(numbers)
	}

	return sumsResult
}

func SumAllTails(numbersToSum ...[]int) []int {
	// append and slice[:] functions return a new slice
	var sumsResult []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumsResult = append(sumsResult, 0)
		} else {
			tail := numbers[1:]
			sumsResult = append(sumsResult, Sum(tail))
		}
	}

	return sumsResult
}
