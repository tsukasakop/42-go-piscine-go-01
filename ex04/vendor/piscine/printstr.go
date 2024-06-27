package piscine

import "ft"

func PrintStr(s string) {
	str := []rune(s)
	for _, i := range str {
		ft.PrintRune(i)
	}
}
