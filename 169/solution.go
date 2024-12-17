package solution

func majorityElement(nums []int) int {
	finder := make(map[int]int)

	for _, num := range nums {
		finder[num]++
		if finder[num] > len(nums)/2 {
			return num
		}
	}

	return 0
}
