package conditional

func IsEven(num int) bool {
	/*
		Check if integer is boolean.

		Params:
			num (int): integer to be checked

		Returns:
			bool: true if integer is even else false
	*/
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

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
