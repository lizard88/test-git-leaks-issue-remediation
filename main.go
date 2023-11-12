package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	apikey := "45321234wefera123"
	dbUsername := "username"
	dbPassword := "user_password"
	dbName := "dbname"
	dbHost := "localhost"
	dbPort := 3306

	fmt.Printf("API key: %s\n",apikey)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connessione al database MySQL riuscita!")

	// Esempio di query di selezione
	rows, err := db.Query("SELECT * FROM nome_tabella")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var col1, col2 string
		if err := rows.Scan(&col1, &col2); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Col1: %s, Col2: %s\n", col1, col2)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
