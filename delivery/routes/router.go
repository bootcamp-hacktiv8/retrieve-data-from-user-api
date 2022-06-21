package routes

import (
	_userRouter "latihan-golang-usamah/delivery/routes/user"
	_userRepository "latihan-golang-usamah/repository/user"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Routes(
	userRepository _userRepository.UserRepositoryInterface,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	mount(router, "/users", _userRouter.UserResource{}.UserRoute(userRepository))

	return router
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
