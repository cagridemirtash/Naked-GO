package conditions

import "fmt"

func WorkshopCondition() {
	first, second, third := 2, 10, 5

	if first > second {
		if second > third {
			PrintStatement(first)
		} else if first > third {
			PrintStatement(first)
		} else {
			PrintStatement(third)
		}
	} else if second > third {
		PrintStatement(second)
	} else {
		PrintStatement(third)
	}
	//3 adet int değişken
	//Ekrana en büyük olanı yazdırınız.
}

func PrintStatement(number int) {
	fmt.Printf("\nThe big number is %d", number)
}
