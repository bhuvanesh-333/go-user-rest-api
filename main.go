package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
}

var users []User

func main() {

	http.HandleFunc("/users", usersHandler)

	fmt.Println("Server running at 8080")
	http.ListenAndServe(":8080", nil)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var u User
		json.NewDecoder(r.Body).Decode(&u)
		users = append(users, u)
		fmt.Fprintln(w, "User Added")

	case http.MethodDelete:
		users = users[:len(users)-1]
		fmt.Fprintln(w, "Last User Deleted")

	default:
		fmt.Fprintln(w, "Method Not Allowed")
	}
}
