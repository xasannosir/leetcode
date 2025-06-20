package solution

type KthLargest struct {
	nums []int
	k    int
}

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

func insert(this *KthLargest, index int, value int) {
	this.nums = append(this.nums[:index], append([]int{value}, this.nums[index:]...)...)
}

func Constructor(k int, nums []int) KthLargest {
	this := KthLargest{
		nums: nums,
		k:    k,
	}

	this.nums = mergeSort(this.nums)

	return this
}

func (this *KthLargest) Add(val int) int {
	left := 0
	right := len(this.nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if val > this.nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	insert(this, left, val)

	return this.nums[len(this.nums)-this.k]
}
