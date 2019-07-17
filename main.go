package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"webapp/routers"

	_ "github.com/lib/pq"
)

// Configuration for webapp service
type Configuration struct {
	WebServicePort int
	DbHost         string
	DbPort         int
	DbName         string
	DbUserName     string
	DbPassword     string
	DbUseSSl       bool
}

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	configuration := Configuration{}
	// Read Config
	file, err := os.Open("config/default.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	initDb(configuration)
	defer db.Close()

	postgresDB := routers.PostgresDB{
		DataMapper: db,
	}

	// Set Router
	http.HandleFunc("/", postgresDB.Router)
	// Start Web Server
	fmt.Println(fmt.Sprintf("appPort: %d", configuration.WebServicePort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.WebServicePort), nil))
}

func initDb(configuration Configuration) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configuration.DbHost, configuration.DbPort,
		configuration.DbUserName, configuration.DbPassword, configuration.DbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

// // Router Public
// func dbRouter(w http.ResponseWriter, r *http.Request) {
// 	routers.Router(w, r, db)
// }
