package main

import "fmt"

func main() {

	items := [5]int{2, 4, 6, 3, 1}

	target := getTargetFromUser()

	x, y := getSumPair(items, target)

	if x != 0 && y != 0 {
		fmt.Printf("(%d,%d)= %d", x, y, target)
	} else {
		fmt.Printf("Aucune pair ne donne comme somme %d", target)
	}

}

func getTargetFromUser() int {
	var target int
	fmt.Println("Saisir un nombre entier naturel")
	fmt.Scan(&target)
	return target
}

func getSumPair(items [5]int, target int) (int, int) {

	var maps = make(map[int]int)
	for i := 0; i < 5; i++ {
		complement := target - items[i]
		if _, isPresent := maps[complement]; isPresent {
			return complement, items[i]
		} else {
			maps[items[i]] = 0 // fictive value
		}
	}
	return 0, 0
}
