package arrays

// Sum takes in a collection of ints and returns its sum
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

//SumAll takes in a collection of int slices and returns a slice where each element
//is the sum of the int slices in the collection in the order given
func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numsToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

//SumAllTails takes in a collection of int slices and returns a slice where
//each element is the sum of all but the first element of the int slices
//in the collection in the order given
func SumAllTails(numsToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
