package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	id    int
	name  string
	email string
}

func users(w http.ResponseWriter, req *http.Request) {

	l := list.New() // Initialize an empty list

	user1 := User{id: 1, name: "aa", email: "aaa@mail"}
	user2 := User{id: 1, name: "aa", email: "aaa@mail"}

	l.PushFront(user1)
	l.PushFront(user2)

	b, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(b))
}

func user(w http.ResponseWriter, req *http.Request) {
	person := User{id: 1, name: "aa", email: "aaa@mail"}

	b, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, string(b))
}

func main() {

	http.HandleFunc("/user", user)
	http.HandleFunc("/user/", users)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
