package main

import (
	"fmt"

	"example/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(quote.Go())
	message := greetings.Hello("xsr")
	fmt.Println(message)
}
