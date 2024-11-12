package solution

func intToRoman(num int) string {
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var roman string

	for i := 0; i < len(numbers); i++ {
		if numbers[i] <= num {
			count := num / numbers[i]
			for j := 0; j < count; j++ {
				roman += romans[i]
			}
			num -= count * numbers[i]
		}
	}

	return roman
}
