package solution

func findMaxK(nums []int) int {
	sort := func(arr []int) {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr)-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}

	sort(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if nums[i] == nums[j]*-1 {
				return nums[i]
			} else if nums[i] > nums[j]*-1 {
				break
			}
		}
	}

	return -1
}
