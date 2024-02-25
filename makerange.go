package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	res := make([]int, max-min)
	count := 0
	for i := min; i < max; i++ {
		res[count] = min + count
		count++
	}

	return res
}
