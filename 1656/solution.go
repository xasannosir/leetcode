package solution

type OrderedStream struct {
	ptr int
	arr []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		ptr: 0,
		arr: make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.arr[idKey-1] = value

	var result []string

	for this.ptr < len(this.arr) && this.arr[this.ptr] != "" {
		result = append(result, this.arr[this.ptr])
		this.ptr++
	}

	return result
}
