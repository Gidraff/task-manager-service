package helpers

import "regexp"


func  IsUserInfoValid(username string, email string, password string) (bool, string) {
	if len(username) < 3 {
		return false, "Invalid username. Should be at least 3 characters long."
	}
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false, "Invalid email format."
	}
	if len(password) < 8 {
		return false, "Invalid password length. Try a longer password"
	}
	return true, "Validation successful"
}
