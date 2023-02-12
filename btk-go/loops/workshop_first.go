package loops

import "fmt"

func GuessingGame() {
	var realNumber int = 12
	var inputFromUser int
	for realNumber != inputFromUser {
		_, err := fmt.Scanln(&inputFromUser)
		if err != nil {
			fmt.Println("This is not number")
			continue
		}
		if realNumber > inputFromUser {
			fmt.Println("Increase your number.")
		}
		if realNumber < inputFromUser {
			fmt.Println("Decrease your number")
		}
	}
	fmt.Println("You won")
}
