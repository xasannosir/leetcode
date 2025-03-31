package solution

func reformatDate(date string) string {
	months := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	var year, month, day string

	if len(date) == 13 {
		day = string(date[0]) + string(date[1])
		month = months[date[5:8]]
		year = date[9:]
	} else {
		day = "0" + string(date[0])
		month = months[date[4:7]]
		year = date[8:]
	}

	return year + "-" + month + "-" + day
}
