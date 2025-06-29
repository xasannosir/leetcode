package solution

type StockSpanner struct {
	nums []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		nums: make([]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	this.nums = append(this.nums, price)
	count := 0

	for i := len(this.nums) - 1; i >= 0; i-- {
		if this.nums[i] <= price {
			count++
		} else {
			break
		}
	}

	return count
}
