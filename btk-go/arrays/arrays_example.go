package arrays

import "fmt"

func ArraysExample() {
	var numbers [5]int
	for i, _ := range numbers {
		numbers[i] = i + 1
	}
	for _, v := range numbers {
		fmt.Println(v)
	}
}
