package solution

type CustomStack struct {
	nums  []int
	count int
	limit int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		nums:  make([]int, 0),
		count: 0,
		limit: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if this.count < this.limit {
		this.nums = append(this.nums, x)
		this.count++
	}
}

func (this *CustomStack) Pop() int {
	if this.count > 0 {
		val := this.nums[this.count-1]
		this.nums = this.nums[:this.count-1]
		this.count--

		return val
	} else {
		return -1
	}
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < this.count; i++ {
		this.nums[i] += val
	}
}
