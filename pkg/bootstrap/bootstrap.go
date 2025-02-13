package bootstrap

import (
	"log"
	"os"
	"server_go/internal/domain"
	"server_go/internal/user"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "Standard output ->", log.LstdFlags|log.Lshortfile)

}

//db es una base de datos. Es un objeto con 2 campos. Un slice de user, y un maxID de inicializacion.

func NewDB() user.DB {
	return user.DB{
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
}
