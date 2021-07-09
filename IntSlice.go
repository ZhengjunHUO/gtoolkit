package gtoolkit

import (
	"fmt"
	"strings"
)

func IntSliceToString(a []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), ""), "[]")	
}

func IntSliceRemove(s []int, idx int) []int {
	return append(s[:idx], s[idx+1:]...)
}
