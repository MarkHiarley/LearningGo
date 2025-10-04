package main

import (
	"encoding/json"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	type pessoa struct {
		Id   int
		Name string
		Age  int
	}

	var pessoas = []pessoa{
		{
			Id:   1,
			Name: "teste,",
			Age:  12,
		},
		{
			Id:   1,
			Name: "teste,",
			Age:  12,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(pessoas)
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}
