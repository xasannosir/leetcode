package solution

type NumArray struct {
	numbers []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		numbers: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	total := 0

	for i := left; i <= right; i++ {
		total += this.numbers[i]
	}

	return total
}
