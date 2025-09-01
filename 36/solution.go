package solution

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if row[board[i][j]] {
				return false
			}
			if board[i][j] != '.' {
				row[board[i][j]] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		col := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if col[board[j][i]] {
				return false
			}
			if board[j][i] != '.' {
				col[board[j][i]] = true
			}
		}
	}

	for i := 0; i <= 6; i += 3 {
		for j := 0; j <= 6; j += 3 {
			box := make(map[byte]bool)
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ {
					if box[board[k][l]] {
						return false
					}
					if board[k][l] != '.' {
						box[board[k][l]] = true
					}
				}
			}
		}
	}

	return true
}
