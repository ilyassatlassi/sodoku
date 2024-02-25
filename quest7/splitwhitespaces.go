package piscine

func SplitWhiteSpaces(s string) []string {
	res := []string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			str += string(s[i])
		} else if s[i] == ' ' {
			if str != "" {
				res = append(res, str)
				str = ""
			}
		}
	}
	if str != "" {
		res = append(res, str)
	}
	return res
}
