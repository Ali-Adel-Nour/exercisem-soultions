package chessboard


type File []bool

// Chessboard contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(board Chessboard, file string) int {
	fileData, exists := board[file]
	if !exists {
		return 0
	}
	count := 0
	for _, isOccupied := range fileData {
		if isOccupied {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard, within the given rank.
func CountInRank(board Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	count := 0
	for _, file := range board {
		// Index is rank - 1 (e.g., Rank 1 is index 0)
		if file[rank-1] {
			count++
		}
	}
	return count
}

// CountAll counts how many squares are present in the chessboard.
func CountAll(board Chessboard) int {
	count := 0
	for _, file := range board {
		// We use 'range' here to count the elements in the slice
		for range file {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(board Chessboard) int {
	count := 0
	for _, file := range board {
		for _, isOccupied := range file {
			if isOccupied {
				count++
			}
		}
	}
	return count
}