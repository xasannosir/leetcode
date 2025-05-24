package solution

func squareIsWhite(coordinates string) bool {
	return (int(rune(coordinates[0]-48))+int(rune(coordinates[1]-48)))%2 == 1
}
