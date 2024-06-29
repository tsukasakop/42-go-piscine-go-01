package piscine

func Atoi(s string) int {
	str := []rune(s)
	nb := 0
	sign := 1
	if str[0] == '-' {
		sign = -1
	}
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}
	for _, c := range str {
		if c < '0' || c > '9' {
			return 0
		}
		nb *= 10
		nb += int(c - '0')
	}
	return sign * nb
}

