package tests

import (
	"testing"

	"github.com/abhinna1/LearnGo/tutorials/conditional"
)

func TestIsEven(t *testing.T) {
	const oddNumber int = 1
	const evenNumber int = 2

	if conditional.IsEven(oddNumber) {
		t.Errorf("Number is odd")
	}

	if !conditional.IsEven(evenNumber) {
		t.Errorf("Number is even.")
	}
}
