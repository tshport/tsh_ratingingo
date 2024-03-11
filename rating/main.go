package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Thank you for engaging our service!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a rating between 1 and 5 for today's service:")

	for {
		// ok or error
		input, _ := reader.ReadString('\n')
		numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if !(numRating < 6 && numRating > 0) {
			fmt.Println("Error input. Please try again by entering a number between 1 and 5.")
			// panic(err)
		} else {
			fmt.Printf("Thanks for giving today's service a rating of %v star(s)!\n", numRating)
			if err != nil {
				fmt.Println(err)
				// panic(err)
			} else {
				fmt.Printf("Added 1 to your rating: %v star(s).", numRating+1)
				break
			}
		}
	}
}
