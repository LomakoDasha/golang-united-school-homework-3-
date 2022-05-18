package homework

func reverse(input []int64) (result []int64) {
	var newSlice []int64
	for i := len(input) - 1; i >= 0; i-- {
		newSlice = append(newSlice, input[i])
	}
	return newSlice
}
