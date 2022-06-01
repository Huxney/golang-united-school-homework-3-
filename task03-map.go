package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, len(input))
	i := 0
	for k := range input {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	for _, k := range keys {
		result = append(result, input[k])
	}

	return result
}
