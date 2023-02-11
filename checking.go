package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*	var input int
		var inputStr string

		for {
			fmt.Print("Enter a number between 1 and 5: ")
			fmt.Scan(&inputStr)
			input, err := strconv.Atoi(strings.TrimSpace(inputStr))
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			if input < 1 || input > 5 {
				fmt.Println("Invalid input. Please enter a number between 1 and 5.")
				continue
			}
			break
		}
		fmt.Println("You entered:", input)    */
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your full name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is:", name)

}
