package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 3 {
		os.Stderr.WriteString(fmt.Sprintf("Program requires two arguments (string, int)).\n"))
		os.Exit(1)
	}

	if age, err := strconv.Atoi(os.Args[2]); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Second parametr must by number.\n"))
		os.Exit(1)
	} else if age <= 18 {
		os.Stderr.WriteString(fmt.Sprintf("Sorry, program is for adults only.\n"))
		os.Exit(0)
	}

	fmt.Printf("Hello %s!\n", os.Args[1])
}
