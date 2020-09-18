package main

import (
	"fmt"
	"math"
)

func modulo() {

	fmt.Print("Write a number: ")
	fmt.Scanln(&firstNumber)

	fmt.Print("Write one more number: ")
	fmt.Scanln(&secondNumber)

	answer := math.Mod(firstNumber, secondNumber)

	fmt.Println("Here is a modulo: ", answer)
	main()
}
