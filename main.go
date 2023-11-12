// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	apikey := "45321234wefera123"
// 	dbUsername := "username"
// 	dbPassword := "user_password"
// 	dbName := "dbname"
// 	dbHost := "localhost"
// 	dbPort := 3306

// 	fmt.Printf("API key: %s\n",apikey)

// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

// 	db, err := sql.Open("mysql", connectionString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connessione al database MySQL riuscita!")

// 	// Esempio di query di selezione
// 	rows, err := db.Query("SELECT * FROM nome_tabella")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var col1, col2 string
// 		if err := rows.Scan(&col1, &col2); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Printf("Col1: %s, Col2: %s\n", col1, col2)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }


package main

import (
	"fmt"
	"os"
)

const (
	// ATTENZIONE: Questa è una chiave segreta hardcoded a scopo illustrativo.
	apiKey = "API_KEY_SEGRETA"
)

func main() {
	// ATTENZIONE: Questo codice stampa la chiave segreta a scopo illustrativo.
	// In un'applicazione reale, non dovresti mai fare ciò con informazioni sensibili.
	fmt.Println("La chiave segreta è:", apiKey)

	// Simuliamo l'uso di una chiave segreta in una richiesta API, ad esempio.
	apiResponse, err := makeAPIRequest(apiKey)
	if err != nil {
		fmt.Println("Errore nella richiesta API:", err)
		os.Exit(1)
	}

	fmt.Println("Risposta API:", apiResponse)
}

// Simula una richiesta API usando la chiave segreta.
func makeAPIRequest(key string) (string, error) {
	// Qui avresti la logica per effettuare una richiesta API autenticata con la chiave segreta.
	// Per scopi illustrativi, questa funzione restituirà solo una stringa di risposta fittizia.
	return "Risposta API fittizia", nil
}
