package solution

type Cashier struct {
	n           int
	discount    int
	products    []int
	prices      []int
	curCustomer int
	productIdxs map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	c := Cashier{
		n:           n,
		discount:    discount,
		products:    products,
		prices:      prices,
		curCustomer: 0,
		productIdxs: map[int]int{},
	}

	for i := 0; i < len(products); i++ {
		c.productIdxs[products[i]] = i
	}

	return c
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.curCustomer++

	var bill float64 = 0

	for i, idx := range product {
		bill += float64(amount[i] * this.prices[this.productIdxs[idx]])
	}

	if this.curCustomer%this.n == 0 {
		bill = bill * ((100.0 - float64(this.discount)) / 100)
	}

	return bill
}
