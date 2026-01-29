package solution

import "slices"

func getLeastFrequentDigit(n int) int {
	freq := make(map[int]int)
	nums := make([]int, 0)
	counts := make([]int, 0)
	digits := make([]int, 0)

	for n != 0 {
		last := n % 10

		freq[last]++

		n /= 10
	}

	for k, v := range freq {
		nums = append(nums, k)
		counts = append(counts, v)
	}

	for i := 0; i < len(counts); i++ {
		for j := 0; j < len(counts)-i-1; j++ {
			if counts[j] > counts[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				counts[j], counts[j+1] = counts[j+1], counts[j]
			}
		}
	}

	leastOccur := counts[0]

	for i := 0; i < len(counts); i++ {
		if counts[i] == leastOccur {
			digits = append(digits, nums[i])
		} else {
			break
		}
	}

	return slices.Min(digits)
}
