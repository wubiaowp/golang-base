package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("hello go !")
	input, error := inputReader.ReadString('\n')
	if error != nil {
		fmt.Printf("error %s\n", error)
		os.Exit(1)
	} else {
		name := input[len(input)-1]
		fmt.Printf("hello %s! what can i do for you", name)
	}

}
