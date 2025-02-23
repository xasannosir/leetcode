package solution

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxStore := 0

	for left < right {
		val := (right - left) * min(height[left], height[right])
		maxStore = max(maxStore, val)

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxStore
}
