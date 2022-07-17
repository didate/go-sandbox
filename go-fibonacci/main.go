package main

import "fmt"

func main() {

	userInput := getUserInput()
	fibonacci := fib(userInput, 0, 1)
	fmt.Println(fibonacci)
}

func fib(n, a, b int) int {

	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else {
		return fib(n-1, b, a+b)
	}

}
func getUserInput() int {
	fmt.Println("Please enter a natural integer")
	var input int
	fmt.Scan(&input)
	return input
}
