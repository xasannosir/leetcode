package solution

func countSymmetricIntegers(low int, high int) int {
	count := 0

	var findLen func(num int, l int) (int, int)
	findLen = func(num int, l int) (int, int) {
		if num > 9 {
			return findLen(num/10, l+1)
		}

		return 0, l + 1
	}

	var sum func(num int) int
	sum = func(num int) int {
		if num > 9 {
			return num%10 + sum(num/10)
		}

		return num
	}

	isEqual := func(num int, lenght int) bool {
		right := 0
		n := 0
		part := lenght / 2

		for num != 0 {
			if n == part {
				break
			}
			n++
			right += num % 10
			num /= 10
		}

		return right == sum(num)
	}

	for i := low; i <= high; i++ {
		_, l := findLen(i, 0)

		if l%2 == 0 && isEqual(i, l) {
			count++
		}
	}

	return count
}
