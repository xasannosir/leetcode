package solution

func destCity(paths [][]string) string {
	finder := func(key string) string {
		for _, path := range paths {
			if path[0] == key {
				return path[1]
			}
		}

		return ""
	}

	curLatest := paths[0][1]

	for {
		latest := finder(curLatest)

		if latest == "" {
			return curLatest
		} else {
			curLatest = latest
		}
	}
}
