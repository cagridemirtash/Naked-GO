package loops

import "fmt"

func FriendsNumber() {
	firstNumber := 220
	secondNumber := 284

	firstSum := 0
	secondSum := 0
	for indexFirst := 1; indexFirst <= firstNumber/2; indexFirst++ {
		if firstNumber%indexFirst == 0 {
			firstSum += indexFirst
		}
	}
	for indexSecond := 1; indexSecond <= secondNumber/2; indexSecond++ {
		if secondNumber%indexSecond == 0 {
			secondSum += indexSecond
		}
	}
	if firstSum == secondNumber && secondSum == firstNumber {
		fmt.Printf("%v and %v are friends number.", firstNumber, secondNumber)
	} else {
		fmt.Printf("%v and %v are not friends number.", firstNumber, secondNumber)
	}
}
