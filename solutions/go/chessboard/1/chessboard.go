package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	output := 0
	for _, val := range cb[rank] {
		if val {
			output++
		}
	}
	return output
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}
	output := 0
	for _, col := range cb {
		if col[file-1] {
			output++
		}
	}
	return output
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	output := 0
	for key := range cb {
		output += len(cb[key])
	}
	return output
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	output := 0
	for key := range cb {
		output += CountInRank(cb, key)
	}
	return output
}
