package solution

func longestMonotonicSubarray(nums []int) int {
	mx := nums[0]
	mn := nums[0]

	mxCount := 1
	mnCount := 1

	res := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > mx {
			mx = nums[i]
			mxCount++

			res = max(res, mnCount)

			mn = nums[i]
			mnCount = 1
		} else if nums[i] < mn {
			mn = nums[i]
			mnCount++

			res = max(res, mxCount)

			mx = nums[i]
			mxCount = 1
		} else {
			res = max(res, max(mnCount, mxCount))
			mx, mn = nums[i], nums[i]
			mxCount, mnCount = 1, 1
		}
	}

	res = max(res, max(mnCount, mxCount))

	return res
}
