package solution

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var arr []int

	for x > 0 {
		arr = append(arr, x%10)
		x /= 10
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}
