package tests

import (
	"testing"

	"github.com/abhinna1/LearnGo/tutorials/loops"
)

// Function to check if two integer lists are equal.
func assertLists(list1 []int, list2 []int) bool {
	if len(list1) != len(list2) {
		return false
	}
	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func TestForLoop(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var count []int = loops.PrintCounter(10)
	if len(list) != len(count) {
		t.Errorf("Lists are not equal.")
	}
	var areListsEqual bool = assertLists(list, count)
	if !areListsEqual {
		t.Errorf("Lists are not equal.")
	}
}
