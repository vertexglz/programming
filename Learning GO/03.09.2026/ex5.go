package main

import "fmt"

type ContactInfo struct {
	Email string
	Phone string
}

type User struct {
	ID       int
	Username string
	ContactInfo
}

func UpdateEmail(u *User, newEmail string) {
	u.Email = newEmail
}

func main() {
	pavluxa := ContactInfo{"pavel.veretennikov@mail.ru", "+79225151581"}
	pavel := User{123, "Pavel", pavluxa}
	UpdateEmail(&pavel, "pavluxagcup@gmail.com")

	fmt.Println(pavel.ContactInfo)
}
