package solution

type SmallestInfiniteSet struct {
	finder map[int]struct{}
	start  int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		finder: make(map[int]struct{}),
		start:  1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.finder) == 0 {
		this.start++
		return this.start - 1
	} else {
		res := 0

		for k, _ := range this.finder {
			res = k
			break
		}

		for k, _ := range this.finder {
			res = min(res, k)
		}

		delete(this.finder, res)

		if this.start > res {
			return res
		} else {
			this.start++
			return this.start - 1
		}
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	this.finder[num] = struct{}{}
}
