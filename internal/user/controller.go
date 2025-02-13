package user

import (
	"context"
	"errors"
	
)

type (
	Controller func(ctx context.Context,request interface{})(interface{},error)

	EndPoints  struct {
		Create Controller
		GetAll Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func MakeEndPoints(ctx context.Context, s Service)EndPoints {
	return EndPoints{
		Create: makeCreateEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
	}
}

func makeGetAllEndpoint(s Service)Controller{
	return func (ctx context.Context, request interface{})(interface{},error){
		users, err := s.GetAll(ctx)
		if err != nil {

			return nil, err
		}
		return users,nil
	}
}




func makeCreateEndpoint(s Service)Controller{
	return func (ctx context.Context, request interface{})(interface{},error){
		req := request.(CreateReq)

	if req.FirstName == "" {
		return nil, errors.New("first name is required")
	}
	if req.LastName == "" {
		return nil, errors.New("Last name is required")
	}
	if req.Email == "" {
		return nil, errors.New("Email is required")
	}

	user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
	if err != nil {
		return nil, err
	}
	return user,nil
	}

}
