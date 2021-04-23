package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UsersContainer struct {
	Users []User `json:"users"`
}

type User struct {
	ID   int
	Name string
}

//func Get(w http.ResponseWriter, r *http.Request) {
//	var users UsersContainer
//	jsonFile, err := os.Open("data.json")
//	if err != nil {
//		panic(err)
//	}
//	defer jsonFile.Close()
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//
//	if err := json.Unmarshal(byteValue, &users); err != nil {
//		// TODO: handle error
//	}
//	if err != nil {
//		// TODO: handle error
//	}
//
//	result := users
//	if err := json.NewEncoder(w).Encode(result); err != nil {
//		// TODO: handle error
//	}
//}

func main() {
	connStr := "user=uraulasevic password=postgres dbname=gotest sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var users []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Println(u.ID, u.Name)
	}

	//	router := mux.NewRouter()
	//	router.HandleFunc("/users", Get).
	//		Methods("GET")
	//	router.HandleFunc("/user/new", Post).
	//		Methods("POST")
	//	routet.HandleFunc("user/update", Update).
	//		Methods("PUT")
	//	router.HandleFunc("user/{id}", Delete).
	//		Methods("DELETE")
	//	fmt.Println("starting server at :8080")
	//	log.Fatal(http.ListenAndServe(":8080", router))
}
