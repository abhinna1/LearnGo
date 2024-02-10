package variables

import (
	"testing"
)

func TestVariables(t *testing.T) {
	// test const variable.
	if CONSTANT1 != "" {
		t.Errorf("Constant did not match")
	}
	if Number_variable != 1 {
		t.Errorf("Variable did not match.")
	}
}
