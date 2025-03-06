package solution

func findMissingAndRepeatedValues(grid [][]int) []int {
	var arr []int

	for i := 0; i < len(grid); i++ {
		arr = append(arr, grid[i]...)
	}

	for i := 0; i < len(arr)-1; i++ {
		flag := true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}

		if flag {
			break
		}
	}

	twice := 0
	missing := 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			twice = arr[i]
		} else if arr[i]+1 != arr[i+1] && missing == 0 {
			missing = arr[i] + 1
		}
	}

	if missing == 0 && arr[0] != 1 {
		missing = 1
	} else if missing == 0 && arr[len(arr)-1] != len(arr) {
		missing = len(arr)
	}

	return []int{twice, missing}
}
