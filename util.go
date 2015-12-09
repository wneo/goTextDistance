package textdistance

// Min returns the minimum number of passed int slices.
func Min(is ...int) int {
	min := 0
	if len(is) > 0 {
		min = is[0]
	}
	for _, v := range is {
		if min > v {
			min = v
		}
	}
	return min
}

// Max returns the maximum number of passed int slices.
func Max(is ...int) int {
	max := 0
	if len(is) > 0 {
		max = is[0]
	}
	for _, v := range is {
		if max < v {
			max = v
		}
	}
	return max
}
