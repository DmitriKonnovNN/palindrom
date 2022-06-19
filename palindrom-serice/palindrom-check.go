package main

import (
	"fmt"
)

const userInitialGreeting = "Herzlich willkommen in unserem Palindrom-Check-Service %v.\n"

var serviceVersion string = "SPAPSHOT 1.0.0"

type UserData struct {
	firstName string
	lastName  string
	email     string
	userInput string
}

func main() {
	greetUser()
	for {
		firstName, lastName, email, inputString := getUsersInput()
		isValidName, isValidEmail, isValidInput := ValidateUserInput(firstName, lastName, email, inputString)

		if isValidName && isValidEmail && isValidInput {
			isPalindrom := searchForPalindrom(inputString)
			var msg string
			switch isPalindrom {
			case true:
				msg = "a palindrom"
			case false:
				msg = "not a palindrom"
			}
			fmt.Printf("Dear %v %v, your text is %v\n", firstName, lastName, msg)
		}
	}

}

func searchForPalindrom(inputText string) (isPalindrom bool) {

	isPalindrom = false
	length := len(inputText)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {

		fmt.Printf("i = %v j = %v \n", inputText[i], inputText[j])
		if inputText[i] != inputText[j] {
			isPalindrom = false
			break
		} else {
			isPalindrom = true
		}
	}
	return isPalindrom
}

func greetUser() {
	fmt.Printf(userInitialGreeting, serviceVersion)
}

func getUsersInput() (string, string, string, string) {
	var firstN string
	var lastN string
	var email string
	var inputString string

	fmt.Println("Gib deinen Vornamen ein ...")
	fmt.Scan(&firstN)
	fmt.Println("Gib deinen Nachname ein ...")
	fmt.Scan(&lastN)
	fmt.Println("Gib deine Emailadresse ein...")
	fmt.Scan(&email)
	fmt.Println("Gib den Text in, um zu überprüfen, ob da ein Palindrom vorliegt.")
	fmt.Scan(&inputString)
	fmt.Println(inputString)

	return firstN, lastN, email, inputString

}
