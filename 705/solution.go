package solution

type MyHashSet struct {
	check map[int]bool
	data  []int
}

func Constructor() MyHashSet {
	return MyHashSet{
		check: map[int]bool{},
		data:  []int{},
	}
}

func (this *MyHashSet) Add(key int) {
	if _, ok := this.check[key]; !ok {
		this.data = append(this.data, key)
		this.check[key] = true
	}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.check, key)
label:
	for i := 0; i < len(this.data); i++ {
		if this.data[i] == key {
			for j := i; j < len(this.data)-1; j++ {
				this.data[i] = this.data[i+1]
			}
			break label
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.check[key]
	return ok
}
