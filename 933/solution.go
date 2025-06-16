package solution

type RecentCounter struct {
	nums []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		nums: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.nums = append(this.nums, t)

	for this.nums[0] < t-3000 {
		this.nums = this.nums[1:]
	}

	return len(this.nums)
}
