package homework

func reverse(input []int64) (result []int64) {
	output := make([]int64, len(input))
	j := 0
	for i := len(input)-1; i >= 0; i-- {
		output[j] = input[i]
		j++
	}

	return output
}
