package loops

import "fmt"

func GuessingGame() {
	realNumber := 12
	var inputFromUser int
	var guessCounter int
	for realNumber != inputFromUser {
		_, err := fmt.Scanln(&inputFromUser)
		if err != nil || inputFromUser >= 100 || inputFromUser <= 0 {
			fmt.Println("This is not a valid number")
			continue
		}
		if realNumber > inputFromUser {
			fmt.Println("Increase your number.")
		}
		if realNumber < inputFromUser {
			fmt.Println("Decrease your number")
		}
		guessCounter += 1
	}
	theMessageOfGuess := ""
	if 1 <= guessCounter && guessCounter <= 3 {
		theMessageOfGuess = "Extreme good"
	} else if guessCounter <= 6 {
		theMessageOfGuess = "Good"
	} else {
		theMessageOfGuess = "Passable"
	}
	fmt.Printf("%v guess is : %s", guessCounter, theMessageOfGuess)
}
