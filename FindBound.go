package gtoolkit

// Return the index of first "target" appear in "slice"
// if "target" not exist, return the index of smallest element which is greater than "target"
func FindLeftBound(slice []int, target int, less func(i, target int) bool) int {
	n := len(slice)
	l, r := 0, n

	for l < r {
		m := l + (r-l)/2
		if less(m, target) {
			l = m + 1
		}else{
			r = m
		}
	}

	return l
}

// Return the index of last "target" appear in "slice"
// if "target" not exist, return the index of biggest element which is smaller than "target"
func FindRightBound(slice []int, target int, more func(i, target int) bool) int {
	n := len(slice)
	l, r := 0, n

	for l < r {
		m := l + (r-l)/2
		if more(m, target) {
			r = m
		}else{
			l = m + 1
		}
	}

	return l - 1
}
