package homework

func average(input [15]float32) (result float32) {
	var arrLength = len(input)
	var arrSum float32 = 0

	for i := 0; i <= arrLength-1; i++ {
		arrSum += input[i]
	}

	return arrSum / float32(arrLength)
}
