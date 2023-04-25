package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		col := []byte{}
		for j := 0; j < 9; j++ {
			col = append(col, board[j][i])
		}
		if !isSelectionValid(col) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		row := []byte{}

		for j := 0; j < 9; j++ {
			row = append(row, board[i][j])
		}
		if !isSelectionValid(row) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			block := []byte{}
			for k := 0; k < 3; k++ {
				smallRow := board[i*3+k][j*3 : j*3+3]
				block = append(block, smallRow...)
			}
			if !isSelectionValid(block) {
				return false
			}
		}
	}
	return true
}

func isSelectionValid(sel []byte) bool {
	bMap := make(map[byte]bool)
	for _, b := range sel {
		if _, ok := bMap[b]; ok && b != '.' {
			return false
		}
		bMap[b] = true
	}
	return true
}
