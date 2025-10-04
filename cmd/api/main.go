package main

import (
	users "API/internal"
	"net/http"
)

func main() {
	http.HandleFunc("/", users.GetUsers)
	http.ListenAndServe(":8080", nil)
}
