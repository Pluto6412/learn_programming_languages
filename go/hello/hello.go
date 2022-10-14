package main

import (
	"fmt"
	"log"

	"example/greetings"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	fmt.Println("Hello world")
	fmt.Println(quote.Go())
	message, err := greetings.Hello("xsr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
