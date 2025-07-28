package arrayandslice

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	sums = make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

func SumAll2(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) int {
	var sumsAllTail int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumsAllTail += 0
		} else {
			numbersTail := numbers[1:]
			sumsAllTail += Sum(numbersTail)
		}
	}
	return sumsAllTail
}
