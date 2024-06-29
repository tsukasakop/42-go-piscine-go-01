package piscine

func BasicAtoi2(s string) int {
	str := []rune(s)
	nb := 0
	for _, c := range str {
		if c < '0' || c > '9' {
			return 0
		}
		nb *= 10
		nb += int(c - '0')
	}
	return nb
}

