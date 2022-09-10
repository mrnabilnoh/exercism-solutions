package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var total int
	ranks, ok := cb[rank]
	if ok {
		for _, v := range ranks {
			if v {
				total++
			}
		}
	}
	return total
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file >= 1 && file <= 8 {
		var total int
		for _, v := range cb {
			// check if current file on Rank is occupied
			if v[file-1] {
				total++
			}
		}
		return total
	}
	return 0
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var total int
	for _, v := range cb {
		total += len(v)
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var total int
	for i := 1; i <= 8; i++ {
		total += CountInFile(cb, i)
	}
	return total
}
