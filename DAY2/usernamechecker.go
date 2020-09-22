package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

const UsernameRegex string = `^@?(\w){1,15}$`

func main() {

	var usernameInput string
	flag.StringVar(&usernameInput, "username", "Gopher", "The GopherFace Username To Check")
	flag.Parse()

	fmt.Println("GopherFace Username Validation Checker")
	fmt.Println("Checking Syntax for username, \"", usernameInput, "\", resulted in: ", CheckUsernameSyntax(usernameInput), "\n")
}

func CheckUsernameSyntax(username string) bool {

	validationResult := false
	result, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = result.MatchString(username)
	return validationResult
}
