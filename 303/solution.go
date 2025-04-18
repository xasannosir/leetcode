package solution

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		arr: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.arr[i]
	}

	return sum
}
