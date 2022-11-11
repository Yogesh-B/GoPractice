package main

import (
	"fmt"

	greet "github.com/yogesh-b/gotutorial/utils"
	q "rsc.io/quote"
)

func main() {
	fmt.Println("What ever let it be the first line of code!")
	str1 := q.Go()
	fmt.Println(str1)

	str2 := greet.Greet("Hacker")
	fmt.Println(str2)
}
