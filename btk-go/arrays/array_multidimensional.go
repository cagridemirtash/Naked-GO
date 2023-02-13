package arrays

import "fmt"

func MultidimensionalArrays() {
	numbers := [5][5]int{}
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			numbers[i][j] = (i + j) * 5
		}
	}
	fmt.Println("numbers", numbers)
}
