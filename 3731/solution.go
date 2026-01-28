package solution

import "sort"

func findMissingElements(nums []int) []int {
	freq := make(map[int]struct{})
	res := make([]int, 0)

	sort.Ints(nums)

	for _, num := range nums {
		freq[num] = struct{}{}
	}

	for i := nums[0]; i < nums[len(nums)-1]; i++ {
		if _, ok := freq[i]; !ok {
			res = append(res, i)
		}
	}

	return res
}
