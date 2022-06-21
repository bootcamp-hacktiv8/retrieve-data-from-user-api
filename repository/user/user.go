package user

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	_entities "latihan-golang-usamah/entities"
	"net/http"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetAllUser() ([]_entities.User, error) {
	res, err := http.Get("https://random-data-api.com/api/users/random_user?size=10")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var user []_entities.User
	errUnmarshal := json.Unmarshal(body, &user)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return user, nil
}
