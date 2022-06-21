package main

import (
	_database "latihan-golang-usamah/database"
	_routes "latihan-golang-usamah/delivery/routes"
	_userRepository "latihan-golang-usamah/repository/user"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db := _database.GetConnection()
	defer db.Close()

	userRepository := _userRepository.NewUserRepository(db)

	router := _routes.Routes(
		userRepository,
	)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
