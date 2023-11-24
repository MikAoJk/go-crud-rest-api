package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func postUser(w http.ResponseWriter, req *http.Request) {
	user := &User{Id: 1, Name: "aa", Email: "aaa@mail"}

	// TODO store incomming req body of a user i Database

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(b))
}

func getUsers(w http.ResponseWriter, req *http.Request) {
	user1 := User{Id: 1, Name: "aa", Email: "aaa@mail"}
	user2 := User{Id: 2, Name: "bb", Email: "bbb@mail"}

	var listUsers []User
	listUsers = append(listUsers, user1)
	listUsers = append(listUsers, user2)

	b, err := json.Marshal(listUsers)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(b))
}

func main() {

	http.HandleFunc("/users", postUser)
	http.HandleFunc("/users/", getUsers)

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
