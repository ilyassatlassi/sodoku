package piscine

func check(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[string]bool)
	for _, char := range base {
		if string(char) == "-" || string(char) == "+" {
			return false
		} else if !seen[string(char)] {
			seen[string(char)] = true
		} else {
			return false
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	baseMap := make(map[byte]int)
	for i := 0; i < len(base); i++ {
		baseMap[base[i]] = i
	}
	res := 0
	factor := 1
	isNegative := true
	if s[0] == '-' {
		s = s[1:]
		isNegative = true
	}
	if check(base) {
		for i := len(s) - 1; i >= 0; i-- {
			power := baseMap[s[i]]
			res += power * factor
			factor *= len(base)
		}
	} else if isNegative {
		res = -res
	}
	return res
}
