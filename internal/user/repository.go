package user

import (
	"context"
	"log"
	"server_go/internal/domain"
)


//Definimos como va a ser la base de datos
//La base de datos, es un slice de users, con un maximo de id hasta el momento.
type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

type (
	//Creamos una interfaz repositorio la cual va a tener 2 metodos. Create y getall
	//ambos reciben el contexto
	//Create recibe un puntero al user que quiere agregar
	//Get all recibe el slice de todos los usuarios.
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
	}

	//Repo tiene la base de datos y un logger
	repo struct {
		db  DB
		log *log.Logger
	}
)


//Inicializa la base de datos, con los usuarios que tengo hasta el momento
//devuelve el puntero al repositorio
func NewRepo(datab DB, l *log.Logger) Repository {
	return &repo{
		db:  datab,
		log: l,
	}
}


//Create, lo primero que hace es aumentar el user id
//guardar en el usuario nuevo el ID que seria el nuevo maximo ID
//Le hace un append para tener el nuevo repositorio 
//Y la primer validacion de repositorio se crea. 
//devuelve por output,  
func (r *repo) Create(ctx context.Context, user *domain.User) error {
	r.db.MaxUserID++
	user.ID = r.db.MaxUserID
	r.db.Users = append(r.db.Users, *user)
	r.log.Println("repository created")
	return nil
}


//get all devuelve un slice con todos los usuarios y un error, en el caso
//que no haya error, se devuelve nil
func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	r.log.Println("repository get all")
	return r.db.Users, nil

}
