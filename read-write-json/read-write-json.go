package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
		[
			{
				"first_name": "Clark",
				"last_name": "Kent",
				"hair_color": "black",
				"has_dog": true
			},
			{
				"first_name": "Bruce",
				"last_name": "Wayne",
				"hair_color": "black",
				"has_dog": false
			}
		]
	`

	// Read json into a struct
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// Write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Lucas"
	m1.LastName = "Genari"
	m1.HairColor = "Black"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Victor"
	m2.LastName = "Genari"
	m2.HairColor = "Black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
