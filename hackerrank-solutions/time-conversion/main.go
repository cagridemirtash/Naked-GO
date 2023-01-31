package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	// Write your code here
	rString := ""
	theSide := s[len(s)-2:]
	//Get hours on string which use for this strconv package
	hours, err := strconv.Atoi(s[0:2])
	if err != nil {
		panic("This is not hours")
	}
	//theSide variable holds PM or AM
	if theSide == "PM" && hours < 12 {
		hours = hours + 12
	}
	if theSide == "AM" && hours == 12 {
		hours = hours - 12
	}
	//Convert to String. Because of return type "string"
	newHours := strconv.Itoa(hours)
	rString = newHours + s[2:len(s)-2]
	//If the  time is AM, Add "0" front of rString
	if theSide == "AM" {
		rString = "0" + rString
	}
	return rString
}

func main() {
	fmt.Println(timeConversion("12:28:12AM"))
}
