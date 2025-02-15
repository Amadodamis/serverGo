package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"server_go/internal/user"
	"server_go/pkg/bootstrap"
	"server_go/pkg/handler"
)

func main() {
	server := http.NewServeMux() //Inicializa el servidor

	//db es una base de datos. Es un objeto con 2 campos. Un slice de user, y un maxID de inicializacion.
	db := bootstrap.NewDB()

	//Crea un logger con un mensaje+flags para que cada accion get/post/put/delete se manifieste con fecha horario.
	logger := bootstrap.NewLogger()


	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()
	handler.NewUserHTTPServer(ctx, server,user.MakeEndPoints(ctx, service))

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))

}
