package solution

func separateDigits(nums []int) []int {
	res := make([]int, 0)

	reverse := func(arr []int) []int {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}

		return arr
	}

	separate := func(num int) []int {
		items := make([]int, 0)

		for num != 0 {
			items = append(items, num%10)
			num /= 10
		}

		return reverse(items)
	}

	for _, num := range nums {
		res = append(res, separate(num)...)
	}

	return res
}
