package gtoolkit

import (
	"fmt"
	"strings"
)

func IntSliceToString(a []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), ""), "[]")	
}

// Return "new" slice with the specified element removed
// The origin slice is no longer safe to use
func IntSliceRemove(s []int, idx int) []int {
	return append(s[:idx], s[idx+1:]...)
}

// Remove the element in place
func IntSliceRemoveInplace(s *[]int, idx int) {
        copy((*s)[idx:], (*s)[idx+1:])
        (*s)[len(*s)-1] = 0
        *s = (*s)[:len(*s)-1]
}
