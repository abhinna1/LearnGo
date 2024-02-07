package conditional


// Check if integer is boolean.
// 	Params:
// 		num (int): integer to be checked
// 	Returns:
// 		bool: true if integer is even else false	
func IsEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

// Function to verify role different roles.
// 	Params:
// 		role(string): role to be verified.
// 	Returns:
// 		int: number based on the role.
func AccessControl(role string) int {
	switch role {
	case "owner":
		return 1
	case "admin":
		return 2
	case "user":
		return 3
	}
	return 0
}

func AccessControlSwitchCases(role string) int {
	if role == "owner" {
		return 1
	} else if role == "admin" {
		return 2
	} else if role == "user" {
		return 3
	}
	return 0
}
