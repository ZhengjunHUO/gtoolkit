package gtoolkit

func Factorial(n uint64) uint64 {
	if n > 0 {
		rslt := n * Factorial(n-1)
		return rslt
	}

	return 1
}
