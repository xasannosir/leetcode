package solution

type NumberContainers struct {
	indexNum   map[int]int
	numIndices map[int][]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		indexNum:   make(map[int]int),
		numIndices: make(map[int][]int),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if num, ok := this.indexNum[index]; ok {
		for ind, val := range this.numIndices[num] {
			if val == index {
				this.numIndices[num] = append(this.numIndices[num][:ind], this.numIndices[num][ind+1:]...)
				if len(this.numIndices[num]) == 0 {
					delete(this.numIndices, num)
				}
				break
			}
		}
	}

	this.numIndices[number] = append(this.numIndices[number], index)
	this.indexNum[index] = number
}

func (this *NumberContainers) Find(number int) int {
	if indices, ok := this.numIndices[number]; ok {
		minimum := indices[0]

		for _, num := range indices {
			if minimum > num {
				minimum = num
			}
		}

		return minimum
	} else {
		return -1
	}
}
