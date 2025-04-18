package solution

type MyStack struct {
	data []int
	size int
}

func Constructor() MyStack {
	return MyStack{
		data: []int{},
		size: 0,
	}
}

func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
	this.size++
}

func (this *MyStack) Pop() int {
	number := this.data[this.size-1]

	this.data = this.data[:this.size-1]
	this.size--

	return number
}

func (this *MyStack) Top() int {
	return this.data[this.size-1]
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}
