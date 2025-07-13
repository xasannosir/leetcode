package solution

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	checker := map[string]struct{}{
		"electronics": struct{}{},
		"grocery":     struct{}{},
		"pharmacy":    struct{}{},
		"restaurant":  struct{}{},
	}

	tempCode := make([]string, 0)
	tempBusinessLine := make([]string, 0)

label1:
	for i := 0; i < len(code); i++ {
		for j := 0; j < len(code[i]); j++ {
			if code[i][j] == '@' {
				continue label1
			}
		}

		if _, ok := checker[businessLine[i]]; ok && isActive[i] {
			tempCode = append(tempCode, code[i])
			tempBusinessLine = append(tempBusinessLine, businessLine[i])
		}
	}

	for i := 0; i < len(tempCode); i++ {
		value := ""
		for j := 0; j < len(tempCode[i]); j++ {
			item := tempCode[i][j]

			if (65 <= item && item <= 90) || (97 <= item && item <= 122) || (48 <= item && item <= 57) || (item == '_') {
				value += string(item)
			}
		}
		tempCode[i] = value
	}

	for i := 0; i < len(tempBusinessLine); i++ {
		for j := 0; j < len(tempBusinessLine)-i-1; j++ {
			if tempBusinessLine[j][0] > tempBusinessLine[j+1][0] {
				tempBusinessLine[j], tempBusinessLine[j+1] = tempBusinessLine[j+1], tempBusinessLine[j]
				tempCode[j], tempCode[j+1] = tempCode[j+1], tempCode[j]
			} else if tempBusinessLine[j][0] == tempBusinessLine[j+1][0] {
				if len(tempCode[j]) > 0 && tempCode[j][0] > tempCode[j+1][0] {
					tempBusinessLine[j], tempBusinessLine[j+1] = tempBusinessLine[j+1], tempBusinessLine[j]
					tempCode[j], tempCode[j+1] = tempCode[j+1], tempCode[j]
					continue
				}

				forEnd := max(len(tempCode[j]), len(tempCode[j+1]))

				for k := 0; k < forEnd; k++ {
					if tempCode[j][k] > tempCode[j+1][k] {
						tempBusinessLine[j], tempBusinessLine[j+1] = tempBusinessLine[j+1], tempBusinessLine[j]
						tempCode[j], tempCode[j+1] = tempCode[j+1], tempCode[j]
						break
					} else if tempCode[j][k] != tempCode[j+1][k] {
						break
					}
				}
			}
		}
	}

	unique := make([]string, 0)

	for _, c := range tempCode {
		if c != "" {
			unique = append(unique, c)
		}
	}

	return unique
}
