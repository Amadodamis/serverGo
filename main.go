package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func main() {
	http.HandleFunc("/users", UserServer)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

var users []User

func init() {
	users = []User{
		{ID: 1,
			FirstName: "Amado",
			LastName:  "Damis",
			Email:     "amadodamis@gmail.com",
		},
		{ID: 2,
			FirstName: "Javier",
			LastName:  "Milei",
			Email:     "JavierMilei@gmail.com",
		},
		{ID: 3,
			FirstName: "Jair",
			LastName:  "Bolsonaro",
			Email:     "JairBolsonaro@gmail.com",
		},
	}
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	var status int

	switch r.Method {
	case http.MethodGet:
		GetAllUser(w)

	case http.MethodPost:
		status = 200
		w.WriteHeader(status)
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, "success in post")

	default:
		status = 404
		w.WriteHeader(status)
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, "not found")
	}
}

func GetAllUser(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, _ := json.Marshal(users)
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data":%s}`, status, value)

}
