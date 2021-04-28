package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int
	Name string
}

type Repo struct {
	conn *sql.DB
}

func (repo *Repo) Get(w http.ResponseWriter, r *http.Request) {
	var users []User

	rows, err := repo.conn.Query("select * from users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	result := users

	if err := json.NewEncoder(w).Encode(result); err != nil {
		//	TODO: handle error
	}
}

func main() {
	connStr := "user=uraulasevic password=postgres dbname=gotest sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	var repo Repo
	repo.conn = db
	defer repo.conn.Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", repo.Get).
		Methods("GET")
		//	router.HandleFunc("/user/new", Post).
		//		Methods("POST")
		//	routet.HandleFunc("user/update", Update).
		//		Methods("PUT")
		//	router.HandleFunc("user/{id}", Delete).
		//		Methods("DELETE")
	fmt.Println("starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
