package piscine

func StrRev(s string) string {
	str := []rune(s)
	l := 0
	for range s {
		l++
	}
	for i:=0; i<l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str)
}
