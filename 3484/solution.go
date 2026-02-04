package solution

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	coordinates map[string]map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		coordinates: make(map[string]map[string]int),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	row := string(cell[0])
	col := cell[1:]

	if this.coordinates[row] == nil {
		this.coordinates[row] = make(map[string]int)
	}

	this.coordinates[row][col] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	row := string(cell[0])
	col := cell[1:]

	if 65 <= cell[0] && cell[0] <= 90 {
		delete(this.coordinates[row], col)
	}
}

func (this *Spreadsheet) GetValue(formula string) int {
	formula = formula[1:]
	parts := strings.Split(formula, "+")

	x := parts[0]
	y := parts[1]

	val1, val2 := 0, 0

	if 65 <= x[0] && x[0] <= 90 {
		val1 = this.coordinates[string(x[0])][x[1:]]
	} else {
		val1, _ = strconv.Atoi(x)
	}

	if 65 <= y[0] && y[0] <= 90 {
		val2 = this.coordinates[string(y[0])][y[1:]]
	} else {
		val2, _ = strconv.Atoi(y)
	}

	return val1 + val2
}
