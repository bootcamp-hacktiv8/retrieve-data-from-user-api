package user

import (
	_entities "latihan-golang-usamah/entities"
)

type UserServiceInterface interface {
	GetAllUser() ([]_entities.User, error)
}
