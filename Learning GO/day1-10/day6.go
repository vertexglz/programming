package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {

	users := map[int]User{
		1: {ID: 1, Name: "Pavel"},
		2: {ID: 2, Name: "Alex"},
	}

	fmt.Println(users[1].Name)

	for id, user := range users {
		fmt.Println(id, user.Name)
	}

	user := users[1]
	user.Name = "Peter"
	users[1] = user

	fmt.Println(users[1].Name)
	delete(users, 2)
	fmt.Println(users)

}
