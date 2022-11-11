package main

import (
	"fmt"

	"github.com/yogesh-b/step-2/models"
)

func main() {
	fmt.Println("Just a random string to test the output.... blah blah!!!")

	u := models.User{
		ID:        2,
		FirstName: "Yogesh",
		LastName:  "Bavishi",
	}

	fmt.Println(u)
}
