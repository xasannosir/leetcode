package solution

import "fmt"

func binary1count(bin int) int {
	binary := ""

	for bin > 0 {
		binary = string(rune(bin%2)+48) + binary
		bin /= 2
	}
	binary = string(rune(bin)) + binary

	count := 0

	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			count++
		}
	}

	return count
}

func readBinaryWatch(turnedOn int) []string {
	var times []string

	for h := 0; h < 12; h++ {
		i := binary1count(h)
		for m := 0; m < 60; m++ {
			j := binary1count(m)
			if turnedOn == i+j {
				times = append(times, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return times
}
