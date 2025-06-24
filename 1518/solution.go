package solution

func numWaterBottles(numBottles int, numExchange int) int {
	total := numBottles
	remainder := 0

	for numBottles > 0 {
		numBottles += remainder
		total += numBottles / numExchange
		remainder = numBottles % numExchange
		numBottles /= numExchange
	}

	return total
}
