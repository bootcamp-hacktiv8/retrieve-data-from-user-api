package user

import "net/http"

type UserHandlerInterface interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	GetAllUserHandler(w http.ResponseWriter, r *http.Request)
}
