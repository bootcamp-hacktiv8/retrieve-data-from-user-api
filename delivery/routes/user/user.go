package user

import (
	_userHandler "latihan-golang-usamah/delivery/handler/user"
	_userRepository "latihan-golang-usamah/repository/user"
	_userService "latihan-golang-usamah/service/user"
	"net/http"

	"github.com/gorilla/mux"
)

type UserResource struct{}

func (ur UserResource) UserRoute(userRepository _userRepository.UserRepositoryInterface) *mux.Router {
	userService := _userService.NewUserUseCase(userRepository)
	userHandler := _userHandler.NewUserHandler(userService)

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(userHandler.GetAllUserHandler)).Methods("GET")
	return router
}
