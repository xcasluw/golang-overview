package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Lucas",
		LastName:    "Genari",
		PhoneNumber: "5511969398424",
		Age:         32,
	}

	log.Println("O usuário é:", user)
}
