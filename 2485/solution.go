package solution

func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}

	var nums []int
	
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
		nums = append(nums, i)
	}

	var s int
	for i := range nums {
		if sum-s == s+i {
			return i
		}
		s+=i
	}

	return -1
}
