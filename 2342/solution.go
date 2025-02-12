package solution

func maximumSum(nums []int) int {
	var findSum func(num int) int

	findSum = func(num int) int {
		if num <= 9 {
			return num
		} else {
			return num%10 + findSum(num/10)
		}
	}

	largest := make(map[int]int)
	maxSum := -1

	for _, current := range nums {
		sum := findSum(current)
		if val, ok := largest[sum]; !ok {
			largest[sum] = current
		} else {
			maxSum = max(maxSum, val+current)
			largest[sum] = max(val, current)
		}
	}

	return maxSum
}
