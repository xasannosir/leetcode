package solution

type NeighborSum struct {
	matrix [][]int
}

func Constructor(grid [][]int) NeighborSum {
	return NeighborSum{
		matrix: grid,
	}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	total := 0

	i, j := this.findIndex(value)

	if i > 0 {
		total += this.matrix[i-1][j]
	}
	if i < len(this.matrix)-1 {
		total += this.matrix[i+1][j]
	}
	if j > 0 {
		total += this.matrix[i][j-1]
	}
	if j < len(this.matrix)-1 {
		total += this.matrix[i][j+1]
	}

	return total
}

func (this *NeighborSum) DiagonalSum(value int) int {
	total := 0

	i, j := this.findIndex(value)

	if i > 0 && j > 0 {
		total += this.matrix[i-1][j-1]
	}
	if i < len(this.matrix)-1 && j > 0 {
		total += this.matrix[i+1][j-1]
	}
	if i > 0 && j < len(this.matrix)-1 {
		total += this.matrix[i-1][j+1]
	}
	if i < len(this.matrix)-1 && j < len(this.matrix)-1 {
		total += this.matrix[i+1][j+1]
	}

	return total
}

func (this *NeighborSum) findIndex(value int) (int, int) {
	for i := 0; i < len(this.matrix); i++ {
		for j := 0; j < len(this.matrix[i]); j++ {
			if this.matrix[i][j] == value {
				return i, j
			}
		}
	}

	return -1, -1
}
