package arraysandslices

func Sum(numbers []int) (total int) {
	for _, number := range numbers {
		total += number
	}

	return
}

func SumAll(arrays ...[]int) []int {
	var totals []int

	for _, array := range arrays {
		totals = append(totals, Sum(array))
	}

	return totals
}

func SumAllTails(arrays ...[]int) []int {
	var totals []int

	for _, array := range arrays {
		if len(array) == 0 {
			totals = append(totals, 0)
			continue
		}
		totals = append(totals, Sum(array[1:]))
	}

	return totals
}
