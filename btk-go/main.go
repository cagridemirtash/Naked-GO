package main

import (
	"fmt"
	"supercharger/functions"
)

func main() {
	//theNameOfUser := variables.Demo1()
	//variables.Demo2(theNameOfUser)

	//You have 1000TL on bank account
	//You want to take 1500 on account -> Fix this with conditional statement
	//takeMoneyUnit := 1000
	//realBalance := 1500
	//conditions.ControlBalance(&takeMoneyUnit, &realBalance)
	//conditions.ControlBalance(&takeMoneyUnit, &realBalance)

	//conditions.WorkshopCondition()

	//loops.ForLoopDemo()
	//loops.GuessingGame()
	//loops.IsPrime()
	//loops.FriendsNumber()
	//arrays.ArraysExample()
	//arrays.MultidimensionalArrays()
	//slices.SlicesIntro()
	//slices.SlicesCopy()

	//Functions Intro Start
	//var sumAdd int = functions.Add(6, 7)
	//fmt.Println(sumAdd)
	//functions.SayHi("Çağrı")
	//Functions Intro End
	sum, subtraction, multiple, division := functions.StoryProblem(10, 6)
	fmt.Println("Sum of Two Number :", sum)
	fmt.Println("Subtraction of Two Number :", subtraction)
	fmt.Println("Multiple of Two Number :", multiple)
	fmt.Println("Division of Two Number :", division)
}
