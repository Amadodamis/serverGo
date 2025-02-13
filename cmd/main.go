package main

import (
	//"encoding/json"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"server_go/internal/domain"
	"server_go/internal/user"
)

func main() {
	server := http.NewServeMux()  //Inicializa el servidor

	//db es una base de datos. Es un objeto con 2 campos. Un slice de user, y un maxID de inicializacion.
	db := user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "Amado",
			LastName:  "Damis",
			Email:     "amadodamis@gmail.com",
		}, {
			ID:        2,
			FirstName: "Javier",
			LastName:  "Milei",
			Email:     "JavierMilei@gmail.com",
		}, {
			ID:        3,
			FirstName: "Jair",
			LastName:  "Bolsonaro",
			Email:     "JairBolsonaro@gmail.com",
		}},
		MaxUserID: 3,
	}

	//Crea un logger con un mensaje+flags para que cada accion get/post/put/delete se manifieste con fecha horario.
	logger := log.New(os.Stdout, "Standard output ->", log.LstdFlags|log.Lshortfile)

	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndPoints(ctx, service))

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))

}

/*
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
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

var users []User

var maxID uint64

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

	maxID = 3
}

func UserServer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetAllUser(w)

	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var u User
		if err := decode.Decode(&u); err != nil {
			MsgResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		PostUser(w, u)

	default:
		InvalidMethod(w)
	}
}

func GetAllUser(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func PostUser(w http.ResponseWriter, data interface{}) {
	user := data.(User)

	if user.FirstName == "" {
		MsgResponse(w, http.StatusBadRequest, "first name is required")
		return
	}
	if user.LastName == "" {
		MsgResponse(w, http.StatusBadRequest, "Last name is required")
		return
	}
	if user.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "Email is required")
		return
	}

	maxID++
	user.ID = maxID
	users = append(users, user)
	DataResponse(w, http.StatusCreated, user)
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d,"message": "method doesn't exist"}`, status)
}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, message)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data":%s}`, status, value)

}




*/
