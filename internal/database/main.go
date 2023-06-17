package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectEstablish() *sql.DB {
	//NHMutW6M89Vu9U8v"
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "127.0.0.1", 5432, "postgres", "password", "checklist")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected!")

	return db

}
