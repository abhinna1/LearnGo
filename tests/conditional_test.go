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

func TestAccessControl(t *testing.T) {
	const ownerRole string = "owner"
	const adminRole string = "admin"
	const userRole string = "user"
	const invalidRole string = "invalid"

	ownerAccess := conditional.AccessControl(ownerRole)
	adminAccess := conditional.AccessControl(adminRole)
	userAccess := conditional.AccessControl(userRole)
	invalidAccess := conditional.AccessControl(invalidRole)

	if ownerAccess != 1 {
		t.Errorf("Invalid Owner Access.")
	}

	if adminAccess != 2 {
		t.Errorf("Invalid admin Access.")
	}

	if userAccess != 3 {
		t.Errorf("Invalid user Access.")
	}

	if invalidAccess != 0 {
		t.Errorf("Invalid invalid Access.")
	}
}

func TestAccessControlSwitchCases(t *testing.T) {
	const ownerRole string = "owner"
	const adminRole string = "admin"
	const userRole string = "user"
	const invalidRole string = "invalid"

	ownerAccess := conditional.AccessControlSwitchCases(ownerRole)
	adminAccess := conditional.AccessControlSwitchCases(adminRole)
	userAccess := conditional.AccessControlSwitchCases(userRole)
	invalidAccess := conditional.AccessControlSwitchCases(invalidRole)

	if ownerAccess != 1 {
		t.Errorf("Invalid Owner Access.")
	}

	if adminAccess != 2 {
		t.Errorf("Invalid admin Access.")
	}

	if userAccess != 3 {
		t.Errorf("Invalid user Access.")
	}

	if invalidAccess != 0 {
		t.Errorf("Invalid invalid Access.")
	}
}
