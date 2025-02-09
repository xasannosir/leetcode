package solution

func thirdMax(nums []int) int {
	unique := make(map[int]int)
	var orders []int

	for _, num := range nums {
		if _, ok := unique[num]; !ok {
			unique[num]++
			orders = append(orders, num)
		}
	}

	for i := 0; i < len(orders); i++ {
		for j := i + 1; j < len(orders); j++ {
			if orders[i] < orders[j] {
				orders[i], orders[j] = orders[j], orders[i]
			}
		}
	}

	l := len(orders)

	if l < 3 {
		return orders[0]
	} else {
		return orders[2]
	}
}
