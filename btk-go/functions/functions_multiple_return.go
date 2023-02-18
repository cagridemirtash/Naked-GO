package functions

func StoryProblem(firstNumber, secondNumber int) (int, int, int, float32) {
	var sum, subtraction, multiple int = firstNumber + secondNumber, firstNumber - secondNumber, firstNumber * secondNumber
	var division float32 = float32(firstNumber / secondNumber)

	return sum, subtraction, multiple, division
}
