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

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	first := mergeSort(arr[:mid])
	second := mergeSort(arr[mid:])

	return merge(first, second)
}

func removeDuplicates(nums []int) int {
	twice := make(map[int]int)

	for _, val := range nums {
		count, ok := twice[val]

		if ok && count < 2 {
			twice[val]++
		} else if !ok {
			twice[val] = 1
		}
	}

	var arr []int

	for k, v := range twice {
		if v == 1 {
			arr = append(arr, k)
		} else {
			arr = append(arr, k)
			arr = append(arr, k)
		}
	}

	arr = mergeSort(arr)

	for i := 0; i < len(arr); i++ {
		nums[i] = arr[i]
	}

	return len(arr)
}
