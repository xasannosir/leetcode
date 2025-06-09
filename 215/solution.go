package solution

func merge(first, second []int) []int {
	var sorted []int
	i, j := 0, 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			sorted = append(sorted, first[i])
			i++
		} else {
			sorted = append(sorted, second[j])
			j++
		}
	}

	sorted = append(sorted, first[i:]...)
	sorted = append(sorted, second[j:]...)

	return sorted
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	first := mergeSort(nums[:mid])
	second := mergeSort(nums[mid:])

	return merge(first, second)
}

func findKthLargest(nums []int, k int) int {
	nums = mergeSort(nums)
	return nums[len(nums)-k]
}
