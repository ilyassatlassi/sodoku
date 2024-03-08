package main

import (
	"os"
)

func Atoi(str string) (int, bool) {
	res := 0
	isnum := true
	isNegative := false
	if str[0] == '-' {
		for i := 1; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				res = res*10 + (int(str[i]) - '0')
				isNegative = true
			} else {
				return 0, false
			}
		}
	} else if str[0] == '+' {
		for i := 1; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				res = res*10 + (int(str[i]) - '0')
			} else {
				return 0, false
			}
		}
	} else {
		for i := 0; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				res = res*10 + (int(str[i]) - '0')
			} else {
				return 0, false
			}
		}
	}
	if isNegative {
		res = -res
	}
	return res, isnum
}

func printnbr(nbr int) {
	if nbr > 9223372036854775805 || nbr < -9223372036854775805 {
		return
	}
	str := ""
	if nbr == 0 {
		os.Stdout.WriteString("0")
	}
	if nbr < 0 {
		os.Stdout.WriteString("-")
		for nbr < 0 {
			remaind := nbr % 10
			if remaind < 0 {
				remaind = -remaind
			}
			str = string(rune('0'+remaind)) + str
			nbr /= 10
		}
		os.Stdout.WriteString(str + "\n")
	} else {
		for nbr > 0 {
			remaind := nbr % 10
			str = string(rune('0'+remaind)) + str
			nbr /= 10
		}
		os.Stdout.WriteString(str + "\n")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 3 {
		num1, valid1 := Atoi(args[0])
		num2, valid2 := Atoi(args[2])
		res := 0
		if num1 > 2147483647 || num2 > 2147483647 {
			return
		}
		if valid1 && valid2 {
			if args[1] == "+" {
				res = num1 + num2
				printnbr(res)
			} else if args[1] == "*" {
				res = num1 * num2
				printnbr(res)
			} else if args[1] == "-" {
				res = num1 - num2
				printnbr(res)
			} else if args[1] == "%" {
				if num2 == 0 {
					os.Stdout.WriteString("No modulo by 0\n")
					return
				} else {
					res = num1 % num2
					printnbr(res)
				}
			} else if args[1] == "/" {
				if num2 == 0 {
					os.Stdout.WriteString("No division by 0\n")
					return
				} else {
					res = num1 / num2
					printnbr(res)
				}
			}
		}
	}
}
