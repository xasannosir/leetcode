package solution

func sumOfEncryptedInt(nums []int) int {
	encrypt := func(x int) int {
		res := x % 10

		for x != 0 {
			res = max(res, x%10)
			x /= 10
		}

		return res
	}

	build := func(num int, count int) int {
		res := 0

		for i := 0; i < count; i++ {
			res = res*10 + num
		}

		return res
	}

	lenght := func(num int) int {
		count := 0

		for num != 0 {
			count++
			num /= 10
		}

		return count
	}

	numbers := make([]int, len(nums))
	total := 0

	for i := 0; i < len(nums); i++ {
		numbers[i] = build(encrypt(nums[i]), lenght(nums[i]))
	}

	for _, number := range numbers {
		total += number
	}

	return total
}
