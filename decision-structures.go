package main

import "log"

func main() {
	// If
	myNumber := 5
	if myNumber >= 5 {
		log.Println("number is", myNumber)
	}

	// Switch
	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("var is set to cat")

	case "dog":
		log.Println("var is set to dog")

	case "fish":
		log.Println("var is set to fish")

	default:
		log.Println("var is set to something else")
	}
}
