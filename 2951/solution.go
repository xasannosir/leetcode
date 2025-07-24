package solution

func findPeaks(mountain []int) []int {
	indices := make([]int, 0)

	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			indices = append(indices, i)
		}
	}

	return indices
}
