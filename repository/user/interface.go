package user

import (
	_entities "latihan-golang-usamah/entities"
)

type UserRepositoryInterface interface {
	GetAllUser() ([]_entities.User, error)
}
