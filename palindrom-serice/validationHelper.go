package main

import "strings"

func ValidateUserInput(fName string, lName string, email string, userString string) (
	isValidName bool, isValidEmail bool, isValidInput bool) {

	isValidName = len(fName) >= 1 && len(lName) >= 1
	isValidEmail = strings.Contains(email, "@")
	isValidInput = true

	return isValidName, isValidEmail, isValidInput
}
