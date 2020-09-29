package main

import (
	"fmt"
	"math"
)

func modulo() {

	firstNumber := getFirstNumber()
	secondNumber := getSecondNumber()

	answer := math.Mod(firstNumber, secondNumber)

	fmt.Println("Here is a modulo: ", answer)
	main()
}
