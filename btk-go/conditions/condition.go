package conditions

import "fmt"

func ControlBalance(takeMoneyUnit, realBalance *int) {
	if *takeMoneyUnit <= *realBalance {
		fmt.Println("This process is successfully.")
		*realBalance -= *takeMoneyUnit
	} else {
		fmt.Printf("Your balance is not enough for %dTL.", *takeMoneyUnit)
	}
}
