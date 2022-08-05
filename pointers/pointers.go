package main

import (
	"log"
)

func main() {
	myString := "Green"

	log.Println("my string is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)
	justLog(myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}

func justLog(s string) {
	log.Println("s is set to", s)
}
