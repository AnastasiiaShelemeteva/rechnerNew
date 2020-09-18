package main

import ( 
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func numbersInput() {
	var firstNumber float64
	var secondNumber float64
	fmt.Print("Write a number: ")
	fmt.Scanln(&firstNumber)

	fmt.Print("Write one more number: ")
	fmt.Scanln(&secondNumber)
	
}

func main() {
	var choice int
	fmt.Print(`Welcome in the calculator: 
	Press (1) to do addition 
	Press (2) to do subtraction
	Press (3) to do multiplication
	Press (4) to do division
	Press (5) to do modulo
	Press (6) to go out of the program
	Please choose one of the options: `)
	fmt.Scanln(&choice)
	input(choice)
}

func input(choice int) {
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

	numbersInput()

	fmt.Println("Here is a addition: ", firstNumber+secondNumber)
	main()
}

func subtraction() {

	fmt.Print("Write a number: ")
	fmt.Scanln(&firstNumber)

	fmt.Print("Write one more number: ")
	fmt.Scanln(&secondNumber)

	fmt.Println("Here is a subtraction: ", firstNumber-secondNumber)
	main()
}

func division() {

	fmt.Print("Write a number: ")
	fmt.Scanln(&firstNumber)

	fmt.Print("Write one more number: ")
	fmt.Scanln(&secondNumber)

	fmt.Println("Here is a division: ", firstNumber/secondNumber)
	main()
}

func multiplication() {

	fmt.Print("Write a number: ")
	fmt.Scanln(&firstNumber)

	fmt.Print("Write one more number: ")
	fmt.Scanln(&secondNumber)

	fmt.Print("Here is a multiplication: ", firstNumber*secondNumber)
	main()
}
