package diffmove

import (
	"log"
	"testing"
)

var testdata = []struct {
	start []int
	end   []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{2, 6, 3, 7, 4}},
}

var deletedata = []struct {
	start  []int
	delete int
	end    []int
}{
	{[]int{1, 2, 3, 4}, 0, []int{2, 3, 4}},
	{[]int{1, 2, 3, 4}, 1, []int{1, 3, 4}},
	{[]int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
	{[]int{1, 2, 3, 4}, 3, []int{1, 2, 3}},
}

func TestRemove(t *testing.T) {
	for _, test := range deletedata {
		remEnd := Remove(test.start, test.delete)
		match := len(remEnd) == len(test.end)
		if match {
			for i := range remEnd {
				if remEnd[i] != test.end[i] {
					match = false
				}
			}
		}
		if !match {
			t.Errorf("Remove failed %v, %v -> %v (%v)", test.start, test.delete, remEnd, test.end)
		}
	}
}

func TestDiffMove(t *testing.T) {
	for _, test := range testdata {
		moves := Diff(test.start, test.end)

		current := make([]int, len(test.start))
		copy(current, test.start)
		for _, move := range moves {
			log.Printf("%v now %v", current, move)
			switch move.Move {
			case "Add":
				current = Insert(current, move.Start, move.Value)
			case "Delete":
				current = Remove(current, move.Start)
			case "Move":
				val := current[move.Start]
				current = Remove(current, move.Start)
				if move.End > move.Start {
					current = Insert(current, move.End, val)
				} else {
					current = Insert(current, move.End, val)
				}
			}
		}

		match := len(test.end) == len(current)
		if match {
			for i := range test.end {
				if test.end[i] != current[i] {
					match = false
				}
			}
		}
		if !match {
			t.Errorf("Mismatch between moves and intention: %v -> %v (should be %v)", test.start, current, test.end)
			t.Errorf("Moves were: %v", moves)
		}
	}
}
