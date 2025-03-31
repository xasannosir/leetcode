package solution

func dayOfYear(date string) int {
	isLeapYear := func(year int) bool {
		return year%400 == 0 || (year%4 == 0 && year%100 != 0)
	}

	year := (int(date[0]-48)*1000 + int(date[1]-48)*100 + int(date[2]-48)*10 + int(date[3]-48))
	month := (int(date[5]-48)*10 + int(date[6]-48))
	day := (int(date[8]-48)*10 + int(date[9]-48))

	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if isLeapYear(year) {
		monthDays[1] = 29
	}

	for i := 0; i < month-1; i++ {
		day += monthDays[i]
	}

	return day
}
