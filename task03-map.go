package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	var sliceOfKeys []int
	var sliceOfSortedValues []string

	for key, _ := range input {
		sliceOfKeys = append(sliceOfKeys, key)
	}

	sort.Ints(sliceOfKeys)

	for i := 0; i <= len(sliceOfKeys)-1; i++ {
		sliceOfSortedValues = append(sliceOfSortedValues, input[i])
	}

	return sliceOfSortedValues
}
