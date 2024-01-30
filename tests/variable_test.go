package tests

import (
	"testing"

	"github.com/abhinna1/LearnGo/tutorials/variables"
)

func TestVariables(t *testing.T) {
	// test const variable.
	if variables.CONSTANT1 != "" {
		t.Errorf("Constant did not match")
	}
	if variables.Number_variable != 1 {
		t.Errorf("Variable did not match.")
	}
}
