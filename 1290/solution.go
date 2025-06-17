package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	bin := ""
	dec := 0

	for {
		bin += string(rune(head.Val + 47))

		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}

	pow := func(x int, y int) int {
		res := 1

		for i := 0; i < y; i++ {
			res *= x
		}

		return res
	}

	for i, j := len(bin)-1, 0; i >= 0; i, j = i-1, j+1 {
		dec += pow(2, j) * int(bin[i]-47)
	}

	return dec
}
