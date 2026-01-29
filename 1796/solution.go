package solution

import "sort"

func secondHighest(s string) int {
	freq := make(map[byte]struct{})
	nums := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if 48 <= s[i] && s[i] <= 57 {
			freq[s[i]] = struct{}{}
		}
	}

	for num, _ := range freq {
		nums = append(nums, int(num)-48)
	}

	sort.Ints(nums)

	if len(nums) > 1 {
		return nums[len(nums)-2]
	} else {
		return -1
	}
}
