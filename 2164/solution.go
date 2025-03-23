package solution

func sortEvenOdd(nums []int) []int {
	var even, odd, res []int

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}

	for i := 0; i < len(even); i++ {
		for j := 0; j < len(even)-i-1; j++ {
			if even[j] > even[j+1] {
				even[j], even[j+1] = even[j+1], even[j]
			}
		}
	}

	for i := 0; i < len(odd); i++ {
		for j := 0; j < len(odd)-i-1; j++ {
			if odd[j] < odd[j+1] {
				odd[j], odd[j+1] = odd[j+1], odd[j]
			}
		}
	}

	evenIndex, oddIndex := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			res = append(res, even[evenIndex])
			evenIndex++
		} else {
			res = append(res, odd[oddIndex])
			oddIndex++
		}
	}

	return res
}
