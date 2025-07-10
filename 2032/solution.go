package solution

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	counter1 := make(map[int]int)
	counter2 := make(map[int]int)
	counter3 := make(map[int]int)
	resCounter := make(map[int]int)
	res := make([]int, 0)

	for _, num := range nums1 {
		counter1[num]++
	}

	for _, num := range nums2 {
		counter2[num]++
	}

	for _, num := range nums3 {
		counter3[num]++
	}

	for k := range counter1 {
		_, ok1 := counter2[k]
		_, ok2 := counter3[k]

		if ok1 || ok2 {
			res = append(res, k)
			resCounter[k]++
		}
	}

	for k := range counter2 {
		_, ok1 := counter1[k]
		_, ok2 := counter3[k]
		_, ok3 := resCounter[k]

		if (ok1 || ok2) && !ok3 {
			res = append(res, k)
			resCounter[k]++
		}
	}

	return res
}
