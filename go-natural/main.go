package main

import (
	"fmt"
)

func main() {
	a, b := getUserInput()
	gcd := gcd(a, b)
	fmt.Printf("GCD (%d,%d) = %d", a, b, gcd)
}

func gcd(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	if diff == 0 {
		return a
	} else if a > b {
		a = diff
	} else {
		b = diff
	}
	return gcd(a, b)
}
func getUserInput() (int, int) {
	var a, b int
	fmt.Println("Saisir deux entiers naturels")
	fmt.Scan(&a, &b)
	return a, b
}
