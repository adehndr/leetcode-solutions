package backtracking

func WordSearch(board [][]byte, word string) bool {
	return wordSearch3(board, word)
}

func wordSearch3(board [][]byte, word string) bool {
	row := len(board)
	column := len(board[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, column)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			globalCounter := 0
			wordCounter := 0
			if helperWordSearch3(board, word, i, j, visited, globalCounter, wordCounter, row, column) {
				return true
			}
		}
	}

	return false
}

func helperWordSearch3(
	board [][]byte, word string, starti int, startj int, visited [][]bool,
	globalCounter int, wordCounter int,
	row int, column int,
) bool {
	if wordCounter == len(word) {
		return true
	}

	if starti < 0 || startj < 0 || starti >= row || startj >= column || visited[starti][startj] ||
		string(board[starti][startj]) != string(word[wordCounter]) {
		return false
	}
	visited[starti][startj] = true
	var result bool

	result = (helperWordSearch3(board, word, starti-1, startj, visited, globalCounter, wordCounter+1, row, column) ||
		helperWordSearch3(board, word, starti+1, startj, visited, globalCounter, wordCounter+1, row, column) ||
		helperWordSearch3(board, word, starti, startj-1, visited, globalCounter, wordCounter+1, row, column) ||
		helperWordSearch3(board, word, starti, startj+1, visited, globalCounter, wordCounter+1, row, column))

	visited[starti][startj] = false
	return result
}
