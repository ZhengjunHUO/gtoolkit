package gtoolkit

func Permutation(s []int) [][]int {
	rslt, n := [][]int{}, len(s)

	var perm func([]int, int)
	perm = func(s []int, m int) {
		if m == 1 {
			elem := make([]int, n)
			copy(elem, s)
			rslt = append(rslt, elem)
			return
		}

		for i := 0; i < m; i ++ {
			perm(s, m - 1)
			if m%2 == 0 {
				s[0], s[m-1] = s[m-1], s[0]
			}else{
				s[i], s[m-1] = s[m-1], s[i]

			}
		}
	}

	perm(s, n)
	return rslt
}
