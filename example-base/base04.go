package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")

	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("an error is: %s\n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("Hello,%s! what can I do for you?\n", name)
	}
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("an error is: %s\n", err)
			continue
		}
		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing":
			fmt.Println("Bye")
			os.Exit(1)
		default:
			fmt.Println("I don't know what do you say")
		}
	}
}
