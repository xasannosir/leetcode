package solution

func tupleSameProduct(nums []int) int {
	mp := make(map[int]int)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			mul := nums[i] * nums[j]
			mp[mul]++
		}
	}

	sum := 0

	for _, v := range mp {
		sum += v * (v - 1) * 4
	}

	return sum
}
