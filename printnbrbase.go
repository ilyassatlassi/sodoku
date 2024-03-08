package piscine

import (
	"fmt"
)

func checkis(bs string) bool {
	for i := 0; i < len(bs); i++ {
		for j := i + 1; j < len(bs)-1; j++ {
			if bs[i] == bs[j] || (bs[i] < 48 || (bs[i] > 57 && bs[i] < 65) || (bs[i] > 90 && bs[i] < 97) || bs[i] > 122) {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) string {
	result := ""
	if checkis(base) {
		if len(base) < 2 {
			fmt.Println("")
		}
		isNegative := false
		if nbr < 0 {
			nbr = nbr * (-1)
			isNegative = true
		}
		n := nbr

		baseLen := len(base)

		for n > 0 {
			index := n % baseLen
			result = string(base[index]) + result
			n = n / baseLen
		}

		if isNegative {
			result = "-" + result
		}
	} else {
		return "NV"
	}
	return result
}
