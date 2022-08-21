package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	AvgGrades, Happiness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d.", u.Name, u.Age, u.Money)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	bob := User{Name: "Bob", Age: 22, Money: -50, AvgGrades: 4.2, Happiness: 0.8}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInfo())
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts/", ContactPage)
	http.ListenAndServe(":8080", nil)
}

func main() {

	handleRequest()
}
