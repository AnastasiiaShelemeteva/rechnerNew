package main

import "fmt"

var choice int

func main() {
	fmt.Print(`Welcome in the calculator: 
	Press (1) to do addition 
	Press (2) to do subtraction
	Press (3) to do multiplication
	Press (4) to do division
	Press (5) to do modulo
	Press (6) to go out of the program
	Please choose one of the options: `)
	fmt.Scanln(&choice)
	input()
}

func getFirstNumber() float64 {
	var firstNumber float64

	fmt.Print("Write a number: ")
	fmt.Scanln(&firstNumber)

	return firstNumber

}

func getSecondNumber() float64 {
	var secondNumber float64

	fmt.Print("Write a number: ")
	fmt.Scanln(&secondNumber)

	return secondNumber

}
func input() {
	switch choice {
	case 1:
		addition()
	case 2:
		subtraction()
	case 3:
		multiplication()
	case 4:
		division()
	case 5:
		modulo()
	case 6:
		fmt.Println("You have been exited out of the program")
	default:
		fmt.Println("This is not a command. Try again")
		main()
	}
}

func addition() {

	firstNumber := getFirstNumber()
	secondNumber := getSecondNumber()
	fmt.Println("Here is a addition: ", firstNumber+secondNumber)
	main()
}

func subtraction() {

	firstNumber := getFirstNumber()
	secondNumber := getSecondNumber()

	fmt.Println("Here is a subtraction: ", firstNumber-secondNumber)
	main()
}

func division() {

	firstNumber := getFirstNumber()
	secondNumber := getSecondNumber()

	fmt.Println("Here is a division: ", firstNumber/secondNumber)
	main()
}

func multiplication() {

	firstNumber := getFirstNumber()
	secondNumber := getSecondNumber()

	fmt.Println("Here is a multiplication: ", firstNumber*secondNumber)
	main()
}
