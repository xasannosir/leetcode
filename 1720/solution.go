package solution

func decode(encoded []int, first int) []int {
	arr := make([]int, 0)
	arr = append(arr, first)

	for i := 0; i < len(encoded); i++ {
		arr = append(arr, encoded[i]^arr[i])
	}

	return arr
}
