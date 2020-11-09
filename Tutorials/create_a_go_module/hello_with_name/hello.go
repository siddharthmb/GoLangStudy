package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Siddharth", "Fatema", "Poorva"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
//	message, err = greetings.Hello("Siddharth")
//	if err == nil {
//		log.Output(message)
//	}
//	fmt.Println(message)
	for _, greet := range message {
		fmt.Println(greet)
	}
}
