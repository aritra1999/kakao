package cmd

import (
	"fmt"
	"log"
	"os/user"
)

func Welcome() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}	

	fmt.Printf("welcome %s\n", currentUser.Username)
}