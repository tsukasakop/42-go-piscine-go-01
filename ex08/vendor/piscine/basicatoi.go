package piscine

func BasicAtoi(s string) int {
	str := []rune(s)
	nb := 1
	l := 0
	for range s {
		l++
	}
	for i:=0; i<l; i++ {
		nb *= int(str[i] - '0')
		nb *= 10
	}
	return nb
}
