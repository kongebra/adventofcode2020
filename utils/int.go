package utils

// FindMax func
func FindMax(list []int) int {
	if (len(list) == 0) {
		return 0
	}

	max := list[0]

	for _, value := range list {
		if value > max {
			max = value
		}
	}

	return max
}

// FindMin func
func FindMin(list []int) int {
	if (len(list) == 0) {
		return 0
	}

	min := list[0]

	for _, value := range list {
		if value < min {
			min = value
		}
	}

	return min
}