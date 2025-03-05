package solution

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	sortedArr := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sortedArr = append(sortedArr, left[i])
			i++
		} else {
			sortedArr = append(sortedArr, right[j])
			j++
		}
	}

	sortedArr = append(sortedArr, left[i:]...)
	sortedArr = append(sortedArr, right[j:]...)

	return sortedArr
}

func sortArray(nums []int) []int {
	return mergeSort(nums)
}
