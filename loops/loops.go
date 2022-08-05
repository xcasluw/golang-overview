package main

import "log"

func main() {
	myNumber := 10
	for i := 0; i <= myNumber; i++ {
		log.Println("i is set to", i)
	}

	// Loops em slices
	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for _, animal := range animals {
		log.Println(animal)
	}

	// Loops em maps (key, value)
	animalsMap := make(map[string]string)
	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Fluffy"
	for animalType, animal := range animalsMap {
		log.Println(animalType, animal)
	}

	// Loops em arrays
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@email.com", 30})
	users = append(users, User{"Jason", "Dex", "jason@email.com", 20})
	users = append(users, User{"Jay", "Hall", "jay@email.com", 40})
	users = append(users, User{"Steve", "Sam", "steve@email.com", 10})

	for _, user := range users {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}
}
