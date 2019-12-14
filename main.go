package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type User struct{
	ID int
	Email string
	Phone string
}

var (
	users []User = []User {
		User {
			ID: 1,
			Email: "exemplo@exemplo.com",
			Phone: "+5504799999999",
		},
		User {
			ID: 2,
			Email: "exemplo2@exemplo.com",
			Phone: "+5504798999999",
		},
	}
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	c, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Erro ao converter users", 500)
		return
	}

	w.Write(c)
}

func main() {
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}