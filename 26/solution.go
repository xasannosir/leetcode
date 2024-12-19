package solution

func removeDuplicates(nums []int) int {
	var slice []int

	slice = append(slice, nums[0])
	for i := 1; i < len(nums); i++ {
		if slice[len(slice)-1] != nums[i] {
			slice = append(slice, nums[i])
		}
	}

	for i := 0; i < len(slice); i++ {
		nums[i] = slice[i]
	}

	nums = nums[:len(nums)-1]

	return len(slice)
}
