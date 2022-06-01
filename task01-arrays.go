package homework

func average(input [15]float32) (result float32) {
	var sum float32
	for _, d := range input {
		sum += d
	}
	result = sum / float32(len(input))

	return result
}
