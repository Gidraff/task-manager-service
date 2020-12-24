package helpers

import (
	"errors"
	"regexp"
	"strings"
)

const regexStr = `^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`

// SignUpValidateInput validate signup
func SignUpValidateInput(username string, email string, password string) (bool, error) {
	if len(strings.TrimSpace(username)) < 3 {
		return false, errors.New("Invalid username. Should be at least 3 characters long")
	}
	emailValid, err := isEmailValid(email)
	if !emailValid {
		return false, err
	}
	passwdValid, err := isPasswordValid(password)
	if !passwdValid {
		return false, err
	}
	return true, nil
}

// AuthValidateInput validate input
func AuthValidateInput(email string, password string) (bool, error) {
	emailValid, err := isEmailValid(email)
	if !emailValid {
		return false, err
	}
	passwdValid, err := isPasswordValid(password)
	if !passwdValid {
		return false, err
	}
	return true, nil
}

func isEmailValid(email string) (bool, error) {
	emailErr := errors.New("Invalid email format")
	if m, _ := regexp.MatchString(regexStr, email); !m {
		return false, emailErr
	}
	return true, nil
}

func isPasswordValid(passwd string) (bool, error) {
	passwdErr := errors.New("Invalid password. should be at least 8 characters long")
	if len(strings.TrimSpace(passwd)) < 6 {
		return false, passwdErr
	}
	return true, nil
}

// ValidateProjectName validate Project name and returns bool type
func ValidateProjectName(name string) (bool, error) {
	projectNameErr := errors.New("Invalid project name")
	if len(strings.TrimSpace(name)) < 3 {
		return false, projectNameErr
	}
	return true, nil
}

// ValidateTaskTitle validate Project name and returns bool type
func ValidateTaskTitle(name string) (bool, error) {
	titleErr := errors.New("invalid task title")
	if len(strings.TrimSpace(name)) < 3 {
		return false, titleErr
	}
	return true, nil
}
