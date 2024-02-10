package conditional

import (
	"testing"
	// "github.com/abhinna1/LearnGo/tutorials/conditional"
)

func TestIsEven(t *testing.T) {
	const oddNumber int = 1
	const evenNumber int = 2

	if IsEven(oddNumber) {
		t.Errorf("Number is odd")
	}

	if !IsEven(evenNumber) {
		t.Errorf("Number is even.")
	}
}

func TestAccessControl(t *testing.T) {
	const ownerRole string = "owner"
	const adminRole string = "admin"
	const userRole string = "user"
	const invalidRole string = "invalid"

	ownerAccess := AccessControl(ownerRole)
	adminAccess := AccessControl(adminRole)
	userAccess := AccessControl(userRole)
	invalidAccess := AccessControl(invalidRole)

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

	ownerAccess := AccessControlSwitchCases(ownerRole)
	adminAccess := AccessControlSwitchCases(adminRole)
	userAccess := AccessControlSwitchCases(userRole)
	invalidAccess := AccessControlSwitchCases(invalidRole)

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
