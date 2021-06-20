package gtoolkit

// build an auxiliary table of same size as s
// the value stored at index i indicates longest proper prefix which is also suffix (for range [0, i])
func GetKmpTable(s string) []int {
	r := []rune(s)
	n := len(r)

	tab := make([]int, n)
	tab[0] = 0
	
	for i:=1; i<n; i++ {
		currMax := tab[i-1]

		for currMax > 0 && r[i] != r[currMax] {
			currMax = tab[currMax-1]
		}
		if r[i] == r[currMax] {
			currMax++
		}

		tab[i] = currMax
	}

	return tab
}
