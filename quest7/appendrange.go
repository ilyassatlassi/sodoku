package piscine

func AppendRange(min, max int) []int {
	res := []int{}
	if min >= max {
		res = nil
	}
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}
