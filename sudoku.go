package main

const size = 9

type Cell struct {
	row, col int
}

func validateSudoku(board [][]int) []Cell {
	var invalidCells []Cell

	// Validação das linhas
	for row := 0; row < size; row++ {
		if hasDuplicates(board[row]) {
			for col := 0; col < size; col++ {
				invalidCells = append(invalidCells, Cell{row, col})
			}
		}
	}

	// Validação das colunas
	for col := 0; col < size; col++ {
		column := make([]int, size)
		for row := 0; row < size; row++ {
			column[row] = board[row][col]
		}
		if hasDuplicates(column) {
			for row := 0; row < size; row++ {
				invalidCells = append(invalidCells, Cell{row, col})
			}
		}
	}

	// Validação das regiões
	for startRow := 0; startRow < size; startRow += 3 {
		for startCol := 0; startCol < size; startCol += 3 {
			region := make([]int, size)
			regionIndex := 0
			for row := startRow; row < startRow+3; row++ {
				for col := startCol; col < startCol+3; col++ {
					region[regionIndex] = board[row][col]
					regionIndex++
				}
			}
			if hasDuplicates(region) {
				for row := startRow; row < startRow+3; row++ {
					for col := startCol; col < startCol+3; col++ {
						invalidCells = append(invalidCells, Cell{row, col})
					}
				}
			}
		}
	}

	return invalidCells
}

func hasDuplicates(arr []int) bool {
	seen := make(map[int]bool)
	for _, num := range arr {
		if num != 0 && seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	// Exemplo de tabuleiro de Sudoku
	board := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	invalidCells
