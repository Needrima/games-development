package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	low, high, count := 0, 100, 1

	guesses := map[int]bool{}

	//scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Guess a number between %d and %d\n", low, high)
	fmt.Println("Enter to continue")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for {
		// binary search tree algorithm
		guess := (low + high) / 2

		if _, ok := guesses[guess]; ok {
			fmt.Println("Number not between 0 and 100")
		} else {
			guesses[guess] = true
		}

		fmt.Println("The number is", guess)

		fmt.Println("(a) too high")
		fmt.Println("(b) too low")
		fmt.Println("(c) that's it")
		scanner.Scan()
		resp := scanner.Text()

		if resp == "a" {
			high = guess - 1
			count++
		} else if resp == "b" {
			low = guess + 1
			count++
		} else if resp == "c" {
			fmt.Printf("I won after %d guesses", count)
			break
		} else {
			fmt.Println("Invalid response")
		}
	}

}
