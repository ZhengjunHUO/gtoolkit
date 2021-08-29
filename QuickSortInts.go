package gtoolkit

func partition(s []int, l, r int) int {
	pivot := s[l]

	for l < r {
		for l < r && s[r] >= pivot {
			r--
		}
		s[l], s[r] = s[r], s[l]
		for l < r && s[l] <= pivot {
			l++
		}
		s[l], s[r] = s[r], s[l]
	}

	return l
}

func quicksort(s []int, l, r int) {
	// == means only 1 element in range, not need to be sorted
	if l >= r {
		return
	}
	
	p := partition(s, l, r)

	quicksort(s, l, p-1)
	quicksort(s, p+1, r)
}

func SortInts(s []int) {
	quicksort(s, 0, len(s)-1)	
}
