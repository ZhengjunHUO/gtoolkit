package gtoolkit

// when target doesn't exist, return the insert position
func searchLeftBnd(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)
	for (l < r) {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		}else{
			r = m
		}
	}

	return l
}

func searchRightBnd(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)
	for (l < r) {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m
		}else{
			l = m + 1
		}
	}

	return l - 1
}
