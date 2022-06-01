package homework

import "fmt"

func average(input [15]float32) (result float32) {
	var total float32 = 15
	for _, value := range input {
		total += value
	}
	fmt.Println(total / float32(len(input)))

	return result
}
