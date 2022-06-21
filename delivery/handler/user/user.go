package user

import (
	"encoding/json"
	"latihan-golang-usamah/delivery/helper"
	_userService "latihan-golang-usamah/service/user"
	"net/http"
)

type UserHandler struct {
	userService _userService.UserServiceInterface
}

func NewUserHandler(userService _userService.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userService.GetAllUser()
	switch {
	case err != nil:
		response, _ := json.Marshal(helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		response, _ := json.Marshal(helper.APIResponseSuccess("success get all user", users))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
