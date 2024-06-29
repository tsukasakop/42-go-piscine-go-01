package piscine

func BasicAtoi(s string) int {
	str := []rune(s)
	nb := 0
	for _, c := range str {
		nb *= 10
		nb += int(c - '0')
	}
	return nb
}
