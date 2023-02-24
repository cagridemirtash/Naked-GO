package main

import (
	"fmt"
	"supercharger/pointers"
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

	/*Multiple Return Function Start*/
	/*sum, subtraction, multiple, division := functions.StoryProblem(10, 6)
	fmt.Println("Sum of Two Number :", sum)
	fmt.Println("Subtraction of Two Number :", subtraction)
	fmt.Println("Multiple of Two Number :", multiple)
	fmt.Println("Division of Two Number :", division)*/
	/*Multiple Return Function End*/

	/*Variadic Function Start*/
	//Use Cases: Cross cutting concerns -> Logging, Caching Mechanism, Authentication dependencies
	/*sum := functions.VariadicFunction(3, 5, 7, 9, 16)
	fmt.Println(sum)*/
	/*Variadic Function End*/
	//maps.Intro()
	/* Example of range use case
	numbersArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for _, v := range numbersArr {
		if v%2 != 0 {
			sum += v
		}
	}
	fmt.Println(sum)*/
	/*Pointer Examples Start*/
	number := 20
	pointers.PointerIntro(&number)
	fmt.Println("Main Numbers :", number)
	/*Pointer Examples End*/
}
