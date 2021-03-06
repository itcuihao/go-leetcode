package leetcode0198

func rob(nums []int) int {
	a, b := 0, 0
	for i, num := range nums {
		if i%2 == 0 {
			a = max(a+num, b)
		} else {
			b = max(a, b+num)
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
