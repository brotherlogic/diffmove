package diffmove

import "log"

// Move An addition, removal or move in an array
type Move struct {
	Move  string
	Start int
	End   int
	Value int
}

// Insert inserts a value into a slice
func Insert(slice []int, index, value int) []int {
	log.Printf("Insert: %v @ %v into %v", value, index, slice)
	slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

// Remove removes a value from a slice
func Remove(slice []int, index int) []int {
	log.Printf("Removing %v from %v", index, slice)
	return append(slice[0:index], slice[index+1:len(slice)]...)
}

//Diff transforms the diff between two lists into a series of moves
func Diff(start []int, end []int) []Move {
	var moves []Move

	newStart := make([]int, len(start), len(start)+1)
	copy(newStart, start)

	// Do any removals first
	removeCount := 0
	for i, startVal := range start {
		found := false
		for _, endVal := range end {
			if startVal == endVal {
				found = true
			}
		}

		if !found {
			moves = append(moves, Move{Move: "Delete", Start: i - removeCount, Value: start[i-removeCount]})
			newStart = Remove(newStart, i-removeCount)
			removeCount++
		}
	}

	// Now do additions
	addCount := 0
	for i, endVal := range end {
		found := false
		for _, startVal := range start {
			if startVal == endVal {
				found = true
			}
		}

		if !found {
			moves = append(moves, Move{Move: "Add", Start: i + addCount, Value: end[i]})
			newStart = Insert(newStart, i+addCount, end[i])
			addCount++
		}
	}

	//Now continually move forward -> back until we reach parity
	for i := range end {
		if end[i] != newStart[i] {
			for j := range end[i:] {
				if newStart[i] == end[i+j] {
					log.Printf("Doing move of %v from %v to %v", newStart[i], i, i+j)
					moves = append(moves, Move{Move: "Move", Start: i, End: i + j, Value: newStart[i]})
					newStart = Remove(newStart, i)
					newStart = Insert(newStart, i+j, end[i+j])
				}
			}
		}
	}

	return moves
}
