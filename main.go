package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your first name: ")
    firstName, _ := reader.ReadString('\n')
    firstName = strings.TrimSpace(firstName)

    fmt.Print("Enter your last name: ")
    lastName, _ := reader.ReadString('\n')
    lastName = strings.TrimSpace(lastName)

    fmt.Printf("You first name is: %s, and your last name is %s.\n", firstName, lastName)
}

