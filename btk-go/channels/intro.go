package channels

func EvenNumbers(evenCn chan int) {
	sumEven := 0
	for i := 0; i <= 10; i += 2 {
		sumEven = sumEven + i
	}
	evenCn <- sumEven
}

func OddNumbers(oddCn chan int) {
	sumOdd := 0
	for i := 1; i < 10; i += 2 {
		sumOdd = sumOdd + i
	}
	oddCn <- sumOdd
}
