package piscine

func SortIntegerTable(table []int) {
	l := 0
	for range table {
		l++
	}
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
