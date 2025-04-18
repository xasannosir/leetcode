package solution

type MyQueue struct {
	data []int
	size int
}

func Constructor() MyQueue {
	return MyQueue{
		data: []int{},
		size: 0,
	}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
	this.size++
}

func (this *MyQueue) Pop() int {
	number := this.data[0]

	this.data = this.data[1:]
	this.size--

	return number
}

func (this *MyQueue) Peek() int {
	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}
