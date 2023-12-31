package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
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

	fmt.Printf("database connection string: %s\n",connectionString)

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


// Funzione per caricare una chiave privata da un file.
func loadPrivateKey(path string) (ssh.Signer, error) {
	keyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("impossibile leggere la chiave privata: %v", err)
	}

	privateKey, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("impossibile analizzare la chiave privata: %v", err)
	}

	return privateKey, nil
}