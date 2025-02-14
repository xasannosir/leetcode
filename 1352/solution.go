package solution

type ProductOfNumbers struct {
	array []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		array: []int{},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	this.array = append(this.array, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	product := 1

	for i := len(this.array) - 1; i >= 0 && k > 0; i, k = i-1, k-1 {
		product *= this.array[i]
	}

	return product
}
