package solution

type MapSum struct {
	data map[string]int
}

func Constructor() MapSum {
	return MapSum{
		data: make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.data[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0

	for k, v := range this.data {
		isEqual := true

		if len(k) < len(prefix) {
			continue
		}

		for i := 0; i < len(k) && i < len(prefix); i++ {
			if k[i] != prefix[i] {
				isEqual = false
				break
			}
		}

		if isEqual {
			sum += v
		}
	}

	return sum
}
