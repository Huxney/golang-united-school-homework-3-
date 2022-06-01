package homework

func reverse(input []int64) (result []int64) {
	inputLen := len(input)
	result = make([]int64, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		result[j] = n
	}

	return result
}
