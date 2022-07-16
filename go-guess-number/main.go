package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	guessedNumber := getRandomNumber()

	fmt.Println("Choose a number from 0 to 10")
	for matching := false; !matching; {
		userNumber := getNumberFromUser()

		matching = isMatching(guessedNumber, userNumber)
	}

}

func isMatching(guessedNumber, userNumber int) bool {
	matching := false
	if userNumber > guessedNumber {
		fmt.Println("Your number is too big, Please choose another one")
	} else if userNumber < guessedNumber {
		fmt.Println("Your number is too small, Please choose another one")
	} else {
		fmt.Println("OHH YES YOU GOT IT - YES")
		matching = true
	}
	return matching
}

func getNumberFromUser() int {
	var input int
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Failed to get number from user", err)
	}
	return input
}

func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % 11
}
