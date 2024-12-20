package solution

func categorizeBox(length int, width int, height int, mass int) string {
	volume := length * width * height

	if (volume >= 1_000_000_000 && mass >= 100) || (mass >= 100) && (length >= 10000 || width >= 10000 || height >= 10000) {
		return "Both"
	} else if volume >= 1_000_000_000 || length >= 10000 || width >= 10000 || height >= 10000 || mass >= 10000 {
		return "Bulky"
	} else if mass >= 100 {
		return "Heavy"
	} else {
		return "Neither"
	}
}
