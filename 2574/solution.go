package solution

func leftRightDifference(nums []int) []int {
	abs := func(number int) int {
		if number < 0 {
			return number * -1
		}

		return number
	}

	var leftSum, rightSum, diff []int
	leftSum, rightSum = append(leftSum, 0), append(rightSum, 0)
	n := len(nums)

	sum := 0
	for i := 0; i < n-1; i++ {
		sum += nums[i]
		leftSum = append(leftSum, sum)
	}

	sum = 0
	for i := n - 1; i > 0; i-- {
		sum += nums[i]
		rightSum = append(rightSum, sum)
	}

	for i, j := 0, len(rightSum)-1; i < j; i, j = i+1, j-1 {
		rightSum[i], rightSum[j] = rightSum[j], rightSum[i]
	}

	for i := 0; i < len(leftSum); i++ {
		diff = append(diff, abs(leftSum[i]-rightSum[i]))
	}

	return diff
}
