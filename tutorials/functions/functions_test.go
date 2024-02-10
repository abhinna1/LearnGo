package functions

import "testing"

func TestSum(t *testing.T) {
	const n1 = 2
	const n2 = 2
	const expected = n1 + n2
	var sum int = getSum(n1, n2)
	if sum != expected {
		t.Errorf("Invalid sum.")
	}
}

func TestHelloWorld(t *testing.T) {
	// Function doesn't return any value.
	helloWorld()
}
