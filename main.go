package main

import (
	"fmt"
	"log"

	"os/user"
)

func welcome() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}	

	fmt.Printf("welcome %s\n", currentUser.Username)
}

func main() {
	welcome();
}