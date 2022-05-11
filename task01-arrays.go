package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	var count float32 = 0;
	
	for _, v := range input {
		if (v != 0) {
			sum += v
			count += 1
		}
	}

	return sum /count
}